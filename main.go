package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/davecgh/go-spew/spew"
	"github.com/gagliardetto/solana-go"

	. "github.com/dave/jennifer/jen"
	bin "github.com/gagliardetto/binary"
	. "github.com/gagliardetto/utilz"
	"github.com/rmntim/solana-anchor-go/sighash"
	"golang.org/x/mod/modfile"
)

const generatedDir = "generated"

// TODO:
// - tests where type has field that is a complex enum (represented as an interface): assign a random concrete value from the possible enum variants.
// - when printing tree, check for len before accessing array indexes.

func main() {
	conf.Encoding = EncodingBorsh
	conf.TypeID = TypeIDAnchor

	filenames := FlagStringArray{}
	flag.Var(&filenames, "src", "Path to source; can use multiple times.")
	flag.StringVar(&conf.DstDir, "dst", generatedDir, "Destination folder")
	flag.StringVar(&conf.Package, "pkg", "", "Set package name to generate, default value is metadata.name of the source IDL.")
	flag.BoolVar(&conf.Debug, "debug", false, "debug mode")
	flag.BoolVar(&conf.RemoveAccountSuffix, "remove-account-suffix", false, "Remove \"Account\" suffix from accessors (if leads to duplication, e.g. \"SetFooAccountAccount\")")

	flag.StringVar((*string)(&conf.Encoding), "codec", string(EncodingBorsh), "Choose codec")
	flag.StringVar((*string)(&conf.TypeID), "type-id", string(TypeIDAnchor), "Choose typeID kind")
	flag.StringVar(&conf.ModPath, "mod", "", "Generate a go.mod file with the necessary dependencies, and this module")
	flag.Parse()

	if err := conf.Validate(); err != nil {
		panic(fmt.Errorf("error while validating config: %w", err))
	}

	var ts time.Time
	if GetConfig().Debug {
		ts = time.Unix(0, 0)
	} else {
		ts = time.Now()
	}
	if len(filenames) == 0 {
		Sfln(
			"[%s] No IDL files provided",
			Red(XMark),
		)
		os.Exit(1)
	}
	{
		exists, err := DirExists(GetConfig().DstDir)
		if err != nil {
			panic(err)
		}
		if !exists {
			MustCreateFolderIfNotExists(GetConfig().DstDir, os.ModePerm)
		}
	}

	callbacks := make([]func(), 0)
	defer func() {
		for _, cb := range callbacks {
			cb()
		}
	}()

	for _, idlFilepath := range filenames {
		Sfln(
			"[%s] Generating client from IDL: %s",
			Shakespeare("+"),
			Shakespeare(idlFilepath),
		)
		idlFile, err := os.Open(idlFilepath)
		if err != nil {
			panic(err)
		}

		dec := json.NewDecoder(idlFile)

		var idl IDL

		err = dec.Decode(&idl)
		if err != nil {
			panic(err)
		}
		{
			if idl.State != nil {
				Sfln(
					"%s idl.State is defined, but generator is not implemented yet.",
					OrangeBG("[?]"),
				)
			}
			//if len(idl.Events) > 0 {
			//	Sfln(
			//		"%s idl.Events is defined, but generator is not implemented yet.",
			//		OrangeBG("[?]"),
			//	)
			//}
			//if len(idl.Errors) > 0 {
			//	Sfln(
			//		"%s idl.Errors is defined, but generator is not implemented yet.",
			//		OrangeBG("[?]"),
			//	)
			//}
			//if len(idl.Constants) > 0 {
			//	Sfln(
			//		"%s idl.Constants is defined, but generator is not implemented yet.",
			//		OrangeBG("[?]"),
			//	)
			//}
		}

		// spew.Dump(idl)

		// Create subfolder for package for generated assets:
		packageAssetFolderName := sighash.ToRustSnakeCase(idl.Metadata.Name)
		var dstDirForFiles string
		if GetConfig().Debug {
			packageAssetFolderPath := path.Join(GetConfig().DstDir, packageAssetFolderName)
			MustCreateFolderIfNotExists(packageAssetFolderPath, os.ModePerm)
			// Create folder for assets generated during this run:
			thisRunAssetFolderName := ToLowerCamel(idl.Metadata.Name) + "_" + ts.Format(FilenameTimeFormat)
			thisRunAssetFolderPath := path.Join(packageAssetFolderPath, thisRunAssetFolderName)
			// Create a new assets folder inside the main assets folder:
			MustCreateFolderIfNotExists(thisRunAssetFolderPath, os.ModePerm)
			dstDirForFiles = thisRunAssetFolderPath
		} else {
			if GetConfig().DstDir == generatedDir {
				dstDirForFiles = filepath.Join(GetConfig().DstDir, packageAssetFolderName)
			} else {
				dstDirForFiles = GetConfig().DstDir
			}
		}
		MustCreateFolderIfNotExists(dstDirForFiles, os.ModePerm)

		files, err := GenerateClientFromProgramIDL(idl)
		if err != nil {
			panic(err)
		}

		{
			mdf := &modfile.File{}
			_ = mdf.AddModuleStmt(GetConfig().ModPath)

			mdf.AddNewRequire("github.com/gagliardetto/solana-go", "v1.5.0", false)
			mdf.AddNewRequire("github.com/fragmetric-labs/solana-binary-go", "v0.8.0", false)
			mdf.AddNewRequire("github.com/gagliardetto/treeout", "v0.1.4", false)
			mdf.AddNewRequire("github.com/gagliardetto/gofuzz", "v1.2.2", false)
			mdf.AddNewRequire("github.com/stretchr/testify", "v1.6.1", false)
			mdf.AddNewRequire("github.com/davecgh/go-spew", "v1.1.1", false)
			mdf.Cleanup()

			//callbacks = append(callbacks, func() {
			//	Ln()
			//	Ln(Bold("Don't forget to import the necessary dependencies!"))
			//	Ln()
			//	for _, v := range mdf.Require {
			//		Sfln(
			//			"	go get %s@%s",
			//			v.Mod.Path,
			//			v.Mod.Version,
			//		)
			//	}
			//	Ln()
			//})

			if GetConfig().ModPath != "" {
				mfBytes, err := mdf.Format()
				if err != nil {
					panic(err)
				}
				gomodFilepath := filepath.Join(dstDirForFiles, "go.mod")
				Sfln(
					"[%s] %s",
					Lime(Checkmark),
					MustAbs(gomodFilepath),
				)
				// Write `go.mod` file:
				err = os.WriteFile(gomodFilepath, mfBytes, 0666)
				if err != nil {
					panic(err)
				}
			}
		}

		for _, file := range files {
			// err := file.Render(os.Stdout)
			// if err != nil {
			// 	panic(err)
			// }

			file.File.HeaderComment("Code generated by https://github.com/gagliardetto/anchor-go. DO NOT EDIT.")
			{
				// Save assets:
				assetFileName := file.Name + ".go"
				assetFilepath := path.Join(dstDirForFiles, assetFileName)

				// Create file:
				goFile, err := os.Create(assetFilepath)
				if err != nil {
					panic(err)
				}
				defer goFile.Close()

				// Write generated code file:
				Sfln(
					"[%s] %s",
					Lime(Checkmark),
					MustAbs(assetFilepath),
				)
				err = file.File.Render(goFile)
				if err != nil {
					panic(err)
				}
			}
		}
	}
}

func FormatSighash(buf []byte) string {
	elems := make([]string, 0)
	for _, v := range buf {
		elems = append(elems, strconv.Itoa(int(v)))
	}

	return "[" + strings.Join(elems, ", ") + "]"
}

func GenerateClientFromProgramIDL(idl IDL) ([]*FileWrapper, error) {
	if idl.Address == "" {
		idl.Address = idl.Metadata.Address
	}

	if GetConfig().Package != "" {
		idl.Metadata.Name = GetConfig().Package
	}

	if err := idl.Validate(); err != nil {
		return nil, err
	}

	// configurable address map
	addresses := make(map[string]string)

	files := make([]*FileWrapper, 0)
	{
		// Create and populate Go file that holds all the basic
		// elements of an instruction client:
		file, err := genProgramBoilerplate(idl)
		if err != nil {
			return nil, err
		}

		file.Add(
			Line().Line().
				Func().Id("DecodeInstructions").
				Params(Id("message").Op("*").Qual(PkgSolanaGo, "Message")).
				Params(
					Id("instructions").Index().Op("*").Id("Instruction"),
					Err().Error(),
				).
				BlockFunc(func(body *Group) {
					body.For(List(Id("_"), Id("ins")).Op(":=").Range().Id("message").Dot("Instructions")).Block(
						Var().Id("programID").Qual(PkgSolanaGo, "PublicKey"),
						If(
							List(Id("programID"), Err()).Op("=").Id("message").Dot("Program").Call(Id("ins").Dot("ProgramIDIndex")),
							Err().Op("!=").Nil(),
						).Block(
							Return(),
						),
						If(Op("!").Id("programID").Dot("Equals").Call(Id("ProgramID"))).Block(
							Continue(),
						),
						Var().Id("accounts").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta"),
						If(
							List(Id("accounts"), Err()).Op("=").Id("ins").Dot("ResolveInstructionAccounts").Call(Id("message")),
							Err().Op("!=").Nil(),
						).Block(
							Return(),
						),
						Var().Id("insDecoded").Op("*").Id("Instruction"),
						If(
							List(Id("insDecoded"), Err()).Op("=").Id("decodeInstruction").Call(Id("accounts"), Id("ins").Dot("Data")),
							Err().Op("!=").Nil(),
						).Block(
							Return(),
						),
						Id("instructions").Op("=").Append(Id("instructions"), Id("insDecoded")),
					)
					body.Return()
				}),
		)

		files = append(files, &FileWrapper{
			Name: "instructions",
			File: file,
		})
	}
	{
		// register complex enums:
		for _, typ := range idl.Types {
			registerComplexEnums(&idl, typ)
		}
		for _, typ := range idl.Accounts {
			registerComplexEnums(&idl, typ)
		}
	}

	defs := make(map[string]IdlTypeDef)
	{
		file := NewGoFile(idl.Metadata.Name, true)
		// Declare types from IDL:
		for _, typ := range idl.Types {
			defs[typ.Name] = typ
			file.Add(genTypeDef(&idl, nil, IdlTypeDef{
				Name: typ.Name,
				Type: typ.Type,
			}))
		}
		files = append(files, &FileWrapper{
			Name: "types",
			File: file,
		})
	}

	{
		file := NewGoFile(idl.Metadata.Name, true)
		// Declare account layouts from IDL:
		for _, acc := range idl.Accounts {
			if _, ok := defs[acc.Name]; ok {
				file.Add(genTypeDef(&idl, acc.Discriminator, IdlTypeDef{
					Name: defs[acc.Name].Name + "Account",
					Type: defs[acc.Name].Type,
				}))
			} else {
				panic(`not implemented - only IDL from ("anchor": ">=0.30.0") is available`)
			}
		}
		files = append(files, &FileWrapper{
			Name: "accounts",
			File: file,
		})
	}

	{
		file := NewGoFile(idl.Metadata.Name, true)

		// Declare account layouts from IDL:
		for _, evt := range idl.Events {
			if _, ok := defs[evt.Name]; ok {
				eventDataTypeName := defs[evt.Name].Name + "EventData"
				file.Add(genTypeDef(&idl, evt.Discriminator, IdlTypeDef{
					Name: eventDataTypeName,
					Type: defs[evt.Name].Type,
				}))
				file.Add(Func().Params(Op("*").Id(eventDataTypeName)).Id("isEventData").Params().Block())
			} else {
				panic(`not implemented - only IDL from ("anchor": ">=0.30.0") is available`)
			}
		}

		file.Add(Empty().Var().Id("eventTypes").Op("=").Map(Index(Lit(8)).Byte()).Qual("reflect", "Type").Values(DictFunc(func(d Dict) {
			for _, evt := range idl.Events {
				if def, ok := defs[evt.Name]; ok {
					d[Id(def.Name+"EventDataDiscriminator")] = Id("reflect.TypeOf(" + def.Name + "EventData{})")
				}
			}
		})))

		file.Add(Empty().Var().Id("eventNames").Op("=").Map(Index(Lit(8)).Byte()).String().Values(DictFunc(func(d Dict) {
			for _, evt := range idl.Events {
				if def, ok := defs[evt.Name]; ok {
					d[Id(def.Name+"EventDataDiscriminator")] = Lit(def.Name)
				}
			}
		})))

		// TODO: refactor it
		// to generate import statements
		file.Add(Empty().Var().Defs(Id("_").Op("*").Qual("strings", "Builder").Op("=").Nil()))
		file.Add(Empty().Var().Defs(Id("_").Op("*").Qual("encoding/base64", "Encoding").Op("=").Nil()))
		file.Add(Empty().Var().Defs(Id("_").Op("*").Qual(PkgDfuseBinary, "Decoder").Op("=").Nil())) // TODO: ..
		file.Add(Empty().Var().Defs(Id("_").Op("*").Qual("github.com/gagliardetto/solana-go/rpc", "GetTransactionResult").Op("=").Nil()))
		file.Add(Empty().Var().Defs(Id("_").Op("*").Qual("github.com/mr-tron/base58", "Alphabet").Op("=").Nil()))

		file.Add(Empty().Type().Id("Event").Struct(
			Id("Name").String(),
			Id("Data").Id("EventData"),
		))

		file.Add(Empty().Type().Id("EventData").Interface(
			Id("UnmarshalWithDecoder").Params(Id("decoder").Op("*").Qual(PkgDfuseBinary, "Decoder")).Error(),
			Id("isEventData").Params(),
		))

		file.Add(Empty().Const().Id("eventLogPrefix").Op("=").Lit("Program data: "))

		file.Add(Empty().Func().Id("DecodeEvents").Params(
			Id("txData").Op("*").Qual("github.com/gagliardetto/solana-go/rpc", "GetTransactionResult"),
			Id("targetProgramId").Qual(PkgSolanaGo, "PublicKey"),
			Id("getAddressTables").Func().Params(
				Id("altAddresses").Index().Qual(PkgSolanaGo, "PublicKey"),
			).Params(
				Id("tables").Map(Qual(PkgSolanaGo, "PublicKey")).Qual(PkgSolanaGo, "PublicKeySlice"),
				Err().Error(),
			),
		).Params(
			Id("evts").Index().Op("*").Id("Event"),
			Err().Error(),
		).BlockFunc(func(g *Group) {
			g.Var().Id("tx").Op("*").Qual(PkgSolanaGo, "Transaction")
			g.If(List(Id("tx"), Err()).Op("=").Id("txData").Dot("Transaction").Dot("GetTransaction").Call(), Err().Op("!=").Nil()).Block(
				Return(),
			)

			g.Id("altAddresses").Op(":=").Make(Index().Qual(PkgSolanaGo, "PublicKey"), Len(Id("tx").Dot("Message").Dot("AddressTableLookups")))
			g.For(List(Id("i"), Id("alt")).Op(":=").Range().Id("tx").Dot("Message").Dot("AddressTableLookups")).Block(
				Id("altAddresses").Index(Id("i")).Op("=").Id("alt").Dot("AccountKey"),
			)
			g.If(Len(Id("altAddresses")).Op(">").Lit(0)).BlockFunc(func(g *Group) {
				g.Var().Id("tables").Map(Qual(PkgSolanaGo, "PublicKey")).Qual(PkgSolanaGo, "PublicKeySlice")
				g.If(List(Id("tables"), Err()).Op("=").Id("getAddressTables").Call(Id("altAddresses")), Err().Op("!=").Nil()).Block(
					Return(),
				)
				g.Id("tx").Dot("Message").Dot("SetAddressTables").Call(Id("tables"))
				g.If(Err().Op("=").Id("tx").Dot("Message").Dot("ResolveLookups").Call(), Err().Op("!=").Nil()).Block(
					Return(),
				)
			})

			g.Var().Id("base64Binaries").Index().Index().Byte()
			g.List(Id("logMessageEventBinaries"), Err()).Op(":=").Id("decodeEventsFromLogMessage").Call(Id("txData").Dot("Meta").Dot("LogMessages"))
			g.If(Err().Op("!=").Nil()).Block(
				Return(),
			)

			g.List(Id("emitedCPIEventBinaries"), Err()).Op(":=").Id("decodeEventsFromEmitCPI").Call(
				Id("txData").Dot("Meta").Dot("InnerInstructions"),
				Id("tx").Dot("Message").Dot("AccountKeys"),
				Id("targetProgramId"),
			)
			g.If(Err().Op("!=").Nil()).Block(
				Return(),
			)

			g.Id("base64Binaries").Op("=").Append(Id("base64Binaries"), Id("logMessageEventBinaries").Op("..."))
			g.Id("base64Binaries").Op("=").Append(Id("base64Binaries"), Id("emitedCPIEventBinaries").Op("..."))
			g.List(Id("evts"), Err()).Op("=").Id("parseEvents").Call(Id("base64Binaries"))
			g.Return()
		}))

		file.Add(Empty().Func().Id("decodeEventsFromLogMessage").Params(
			Id("logMessages").Index().String(),
		).Params(
			Id("eventBinaries").Index().Index().Byte(),
			Err().Error(),
		).BlockFunc(func(g *Group) {
			g.For(List(Id("_"), Id("log")).Op(":=").Range().Id("logMessages")).Block(
				If(Qual("strings", "HasPrefix").Call(Id("log"), Id("eventLogPrefix"))).BlockFunc(func(g *Group) {
					g.Id("eventBase64").Op(":=").Id("log").Index(Len(Id("eventLogPrefix")).Op(":"))

					g.Var().Id("eventBinary").Index().Byte()
					g.If(
						List(Id("eventBinary"), Err()).Op("=").Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(Id("eventBase64")),
						Err().Op("!=").Nil(),
					).Block(
						Err().Op("=").Qual("fmt", "Errorf").Call(Lit("failed to decode logMessage event: %s"), Id("eventBase64")),
						Return(),
					)
					g.Id("eventBinaries").Op("=").Append(Id("eventBinaries"), Id("eventBinary"))
				}),
			)
			g.Return()
		}))

		file.Add(Empty().Func().Id("decodeEventsFromEmitCPI").Params(
			Id("InnerInstructions").Index().Qual("github.com/gagliardetto/solana-go/rpc", "InnerInstruction"),
			Id("accountKeys").Qual(PkgSolanaGo, "PublicKeySlice"),
			Id("targetProgramId").Qual(PkgSolanaGo, "PublicKey"),
		).Params(
			Id("eventBinaries").Index().Index().Byte(),
			Err().Error(),
		).BlockFunc(func(g *Group) {
			g.For(List(Id("_"), Id("parsedIx")).Op(":=").Range().Id("InnerInstructions")).Block(
				For(List(Id("_"), Id("ix")).Op(":=").Range().Id("parsedIx").Dot("Instructions")).BlockFunc(func(g *Group) {
					g.If(Id("accountKeys").Index(Id("ix").Dot("ProgramIDIndex")).Op("!=").Id("targetProgramId")).Block(
						Continue(),
					)

					g.Var().Id("ixData").Index().Byte()
					g.If(
						List(Id("ixData"), Err()).Op("=").Qual("github.com/mr-tron/base58", "Decode").Call(Id("ix").Dot("Data").Dot("String").Call()),
						Err().Op("!=").Nil(),
					).Block(
						Return(),
					)
					g.Id("eventBase64").Op(":=").Qual("encoding/base64", "StdEncoding").Dot("EncodeToString").Call(Id("ixData").Index(Lit(8).Op(":")))
					g.Var().Id("eventBinary").Index().Byte()
					g.If(
						List(Id("eventBinary"), Err()).Op("=").Qual("encoding/base64", "StdEncoding").Dot("DecodeString").Call(Id("eventBase64")),
						Err().Op("!=").Nil(),
					).Block(
						Return(),
					)
					g.Id("eventBinaries").Op("=").Append(Id("eventBinaries"), Id("eventBinary"))
				}),
			)
			g.Return()
		}))

		file.Add(Empty().Func().Id("parseEvents").Params(
			Id("base64Binaries").Index().Index().Byte(),
		).Params(
			Id("evts").Index().Op("*").Id("Event"),
			Err().Error(),
		).BlockFunc(func(g *Group) {
			g.Id("decoder").Op(":=").Qual(PkgDfuseBinary, "NewDecoderWithEncoding").Call(Nil(), Qual(PkgDfuseBinary, "EncodingBorsh"))

			g.For(List(Id("_"), Id("eventBinary")).Op(":=").Range().Id("base64Binaries")).BlockFunc(func(g *Group) {
				g.Id("eventDiscriminator").Op(":=").Qual(PkgDfuseBinary, "TypeID").Call(Id("eventBinary").Index(Lit(0).Op(":").Lit(8)))
				g.If(List(Id("eventType"), Id("ok")).Op(":=").Id("eventTypes").Index(Id("eventDiscriminator")), Id("ok")).BlockFunc(func(g *Group) {
					g.Id("eventData").Op(":=").Qual("reflect", "New").Call(Id("eventType")).Dot("Interface").Call().Assert(Id("EventData"))
					g.Id("decoder").Dot("Reset").Call(Id("eventBinary"))
					g.If(Err().Op("=").Id("eventData").Dot("UnmarshalWithDecoder").Call(Id("decoder")), Err().Op("!=").Nil()).Block(
						Err().Op("=").Qual("fmt", "Errorf").Call(Lit("failed to unmarshal event %s: %w"), Id("eventType").Dot("String").Call(), Err()),
						Return(),
					)
					g.Id("evts").Op("=").Append(Id("evts"), Op("&").Id("Event").Values(Dict{
						Id("Name"): Id("eventNames").Index(Id("eventDiscriminator")),
						Id("Data"): Id("eventData"),
					}))
				})
			})
			g.Return()
		}))

		files = append(files, &FileWrapper{
			Name: "events",
			File: file,
		})
	}

	{
		// define custom errors
		file := NewGoFile(idl.Metadata.Name, true)

		// to generate import statements
		file.Add(Var().Defs(
			Id("_").Op("*").Qual("encoding/json", "Encoder").Op("=").Nil(),
			Id("_").Op("*").Qual("github.com/gagliardetto/solana-go/rpc/jsonrpc", "RPCError").Op("=").Nil(),
			Id("_").Qual("fmt", "Formatter").Op("=").Nil(),
			Id("_").Op("=").Qual("errors", "ErrUnsupported"),
		))

		file.Add(Var().DefsFunc(func(group *Group) {
			errDict := Dict{}
			for _, errDef := range idl.Errors {
				name := "Err" + ToCamel(errDef.Name)
				group.Add(Id(name).Op("=").Op("&").Id("customErrorDef").Values(Dict{
					Id("code"): Lit(errDef.Code),
					Id("name"): Lit(errDef.Name),
					Id("msg"):  Lit(errDef.Msg),
				}))
				errDict[Lit(errDef.Code)] = Id(name)
			}
			group.Add(Id("Errors").Op("=").Map(Int()).Id("CustomError").Values(errDict))
		}))

		file.Add(
			Type().Id("CustomError").Interface(
				Id("Code").Params().Int(),
				Id("Name").Params().String(),
				Id("Error").Params().String(),
			),
			Line(),
			Type().Id("customErrorDef").Struct(
				Id("code").Int(),
				Id("name").String(),
				Id("msg").String(),
			),
			Line(),
			Func().Params(Id("e").Op("*").Id("customErrorDef")).Id("Code").Params().Int().Block(
				Return(Id("e").Dot("code")),
			),
			Line(),
			Func().Params(Id("e").Op("*").Id("customErrorDef")).Id("Name").Params().String().Block(
				Return(Id("e").Dot("name")),
			),
			Line(),
			Func().Params(Id("e").Op("*").Id("customErrorDef")).Id("Error").Params().String().Block(
				Return(Qual("fmt", "Sprintf").Call(Lit("%s(%d): %s"), Id("e").Dot("name"), Id("e").Dot("code"), Id("e").Dot("msg"))),
			),
			Line(),
			Func().Id("DecodeCustomError").Params(Id("rpcErr").Error()).Params(Err().Error(), Id("ok").Bool()).Block(
				If(List(Id("errCode"), Id("o")).Op(":=").Id("decodeErrorCode").Call(Id("rpcErr")), Id("o")).Block(
					If(List(Id("customErr"), Id("o")).Op(":=").Id("Errors").Index(Id("errCode")), Id("o")).Block(
						Id("err").Op("=").Id("customErr"),
						Id("ok").Op("=").True(),
						Return(),
					),
				),
				Return(),
			),
			Line(),
			Func().Id("decodeErrorCode").Params(Id("rpcErr").Error()).Params(Id("errorCode").Int(), Id("ok").Bool()).Block(
				Var().Id("jErr").Op("*").Qual("github.com/gagliardetto/solana-go/rpc/jsonrpc", "RPCError"),
				If(Qual("errors", "As").Call(Id("rpcErr"), Op("&").Id("jErr")).Op("&&").Id("jErr").Dot("Data").Op("!=").Nil()).Block(
					If(List(Id("root"), Id("o")).Op(":=").Id("jErr").Dot("Data").Assert(Map(String()).Interface()), Id("o")).Block(
						If(List(Id("rootErr"), Id("o")).Op(":=").Id("root").Index(Lit("err")).Assert(Map(String()).Interface()), Id("o")).Block(
							If(List(Id("rootErrInstructionError"), Id("o")).Op(":=").Id("rootErr").Index(Lit("InstructionError")), Id("o")).Block(
								If(List(Id("rootErrInstructionErrorItems"), Id("o")).Op(":=").Id("rootErrInstructionError").Assert(Index().Interface()), Id("o")).Block(
									If(Len(Id("rootErrInstructionErrorItems")).Op("==").Lit(2)).Block(
										If(List(Id("v"), Id("o")).Op(":=").Id("rootErrInstructionErrorItems").Index(Lit(1)).Assert(Map(String()).Interface()), Id("o")).Block(
											If(List(Id("v2"), Id("o")).Op(":=").Id("v").Index(Lit("Custom")).Assert(Qual("encoding/json", "Number")), Id("o")).Block(
												If(List(Id("code"), Err()).Op(":=").Id("v2").Dot("Int64").Call(), Err().Op("==").Nil()).Block(
													Id("ok").Op("=").True(),
													Id("errorCode").Op("=").Int().Call(Id("code")),
												),
											).Else().If(List(Id("v2"), Id("o")).Op(":=").Id("v").Index(Lit("Custom")).Assert(Float64()), Id("o")).Block(
												Id("ok").Op("=").True(),
												Id("errorCode").Op("=").Int().Call(Id("v2")),
											),
										),
									),
								),
							),
						),
					),
				),
				Return(),
			),
		)

		files = append(files, &FileWrapper{
			Name: "errors",
			File: file,
		})
	}

	// Instructions:
	for _, instruction := range idl.Instructions {
		file := NewGoFile(idl.Metadata.Name, true)
		insExportedName := ToCamel(instruction.Name)
		var args []IdlField
		for _, arg := range instruction.Args {
			idlFieldArg := IdlField{
				Name: arg.Name,
				Docs: arg.Docs,
				Type: IdlType{
					asString:         arg.Type.asString,
					asIdlTypeVec:     arg.Type.asIdlTypeVec,
					asIdlTypeOption:  arg.Type.asIdlTypeOption,
					asIdlTypeArray:   arg.Type.asIdlTypeArray,
					asIdlTypeDefined: nil,
				},
			}
			if arg.Type.asIdlTypeDefined != nil {
				idlFieldArg.Type.asIdlTypeDefined = &IdlTypeDefined{
					Defined: IdLTypeDefinedName{
						Name: arg.Type.asIdlTypeDefined.Defined.Name,
					},
				}
			}
			args = append(args, idlFieldArg)
		}

		// fmt.Println(RedBG(instruction.Name))

		{
			code := Empty().Line().Line()

			for _, doc := range instruction.Docs {
				code.Comment(doc).Line()
			}

			if len(instruction.Docs) == 0 {
				code.Commentf(
					"%s is the `%s` instruction.",
					insExportedName,
					instruction.Name,
				).Line()
			}

			code.Type().Id(insExportedName).StructFunc(func(fieldsGroup *Group) {
				for argIndex, arg := range args {
					if len(arg.Docs) > 0 {
						if argIndex > 0 {
							fieldsGroup.Line()
						}
						for _, doc := range arg.Docs {
							fieldsGroup.Comment(doc)
						}
					}
					fieldsGroup.Add(genField(arg, true)).
						Add(func() Code {
							if arg.Type.IsIdlTypeOption() {
								return Tag(map[string]string{
									"bin": "optional",
								})
							}
							return nil
						}())
				}

				fieldsGroup.Line()

				{
					lastGroupName := ""
					// Add comments of the accounts from rust docs.
					instruction.Accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {
						comment := &strings.Builder{}
						indent := 6
						var prepend int

						if groupPath != "" {
							thisGroupName := filepath.Base(groupPath)
							indent = len(thisGroupName) + 2
							if strings.Count(groupPath, "/") == 0 {
								prepend = 6
							} else {
								prepend = 6 + (strings.Count(groupPath, "/") * 2) + len(strings.TrimSuffix(groupPath, thisGroupName)) - 1
							}
							if lastGroupName != groupPath {
								comment.WriteString(strings.Repeat("·", prepend-1) + Sf(" %s: ", thisGroupName))
							} else {
								comment.WriteString(strings.Repeat("·", prepend+indent-1) + " ")
							}
							lastGroupName = groupPath
						}

						comment.WriteString(Sf("[%v] = ", accountIndex))
						comment.WriteString("[")
						if ia.Writable {
							comment.WriteString("WRITE")
						}
						if ia.Signer {
							if ia.Writable {
								comment.WriteString(", ")
							}
							comment.WriteString("SIGNER")
						}
						comment.WriteString("] ")
						comment.WriteString(ia.Name)

						fieldsGroup.Comment(comment.String())
						for _, doc := range ia.Docs {
							fieldsGroup.Comment(strings.Repeat("·", prepend+indent-1+6) + " " + doc)
						}
						if accountIndex < instruction.Accounts.NumAccounts()-1 {
							fieldsGroup.Comment("")
						}

						return true
					})
				}
				fieldsGroup.Qual(PkgSolanaGo, "AccountMetaSlice").Tag(map[string]string{
					"bin": "-",
				})
			})

			file.Add(code.Line())
		}

		{
			builderFuncName := formatBuilderFuncName(insExportedName)
			code := Empty()
			code.Commentf(
				"%s creates a new `%s` instruction builder.",
				builderFuncName,
				insExportedName,
			).Line()
			//
			code.Func().Id(builderFuncName).Params().Op("*").Id(insExportedName).
				BlockFunc(func(body *Group) {
					body.Id("nd").Op(":=").Op("&").Id(insExportedName).Block(
						Id("AccountMetaSlice").Op(":").Make(Qual(PkgSolanaGo, "AccountMetaSlice"), Lit(instruction.Accounts.NumAccounts())).Op(","),
					)

					// Set sysvar accounts and constant accounts:
					instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
						if isVar(account.Name) {
							pureVarName := getSysVarName(account.Name)
							is := isSysVar(pureVarName)
							if is {
								_, ok := sysVars[pureVarName]
								if !ok {
									panic(account)
								}
								def := Qual(PkgSolanaGo, "Meta").Call(Qual(PkgSolanaGo, pureVarName))
								if account.Writable {
									def.Dot("WRITE").Call()
								}
								if account.Signer {
									def.Dot("SIGNER").Call()
								}
								body.Id("nd").Dot("AccountMetaSlice").Index(Lit(index)).Op("=").Add(def)
							} else {
								panic(account)
							}
						} else if account.Address != "" {
							//def := Qual(PkgSolanaGo, "Meta").Call(Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(account.Address)))
							def := Qual(PkgSolanaGo, "Meta").Call(Id("Addresses").Index(Lit(account.Address)))
							addresses[account.Address] = account.Address
							if account.Writable {
								def.Dot("WRITE").Call()
							}
							if account.Signer {
								def.Dot("SIGNER").Call()
							}
							body.Id("nd").Dot("AccountMetaSlice").Index(Lit(index)).Op("=").Add(def)
						}
						return true
					})

					body.Return(Id("nd"))
				})
			file.Add(code.Line())
		}

		{
			// Declare methods that set the parameters of the instruction:
			code := Empty()
			for _, arg := range args {
				exportedArgName := ToCamel(arg.Name)

				code.Line().Line()
				name := "Set" + exportedArgName
				code.Commentf("%s sets the %q parameter.", name, arg.Name).Line()
				for _, doc := range arg.Docs {
					code.Comment(doc).Line()
				}

				code.Func().Params(Id("inst").Op("*").Id(insExportedName)).Id(name).
					Params(
						ListFunc(func(params *Group) {
							// Parameters:
							params.Id(arg.Name).Add(genTypeName(arg.Type))
						}),
					).
					Params(
						ListFunc(func(results *Group) {
							// Results:
							results.Op("*").Id(insExportedName)
						}),
					).
					BlockFunc(func(body *Group) {
						// Body:
						body.Id("inst").Dot(exportedArgName).Op("=").
							Add(func() Code {
								if isComplexEnum(arg.Type) {
									return Op("&")
								}
								return Op("&")
							}()).
							Id(arg.Name)

						body.Return().Id("inst")
					})
			}

			file.Add(code.Line())
		}
		{
			// Declare methods that set/get accounts for the instruction:
			code := Empty()

			declaredReceivers := []string{}
			groupMemberIndex := 0
			instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
				builderStructName := insExportedName + ToCamel(parentGroupPath) + "AccountsBuilder"
				hasNestedParent := parentGroupPath != ""
				isDeclaredReceiver := SliceContains(declaredReceivers, parentGroupPath)

				if !hasNestedParent {
					groupMemberIndex = index
				}
				if hasNestedParent && !isDeclaredReceiver {
					groupMemberIndex = 0
					declaredReceivers = append(declaredReceivers, parentGroupPath)
					// many accounts (???)
					// builder struct for this accounts group:

					code.Line().Line()
					for _, doc := range parentGroup.Docs {
						code.Comment(doc).Line()
					}
					code.Type().Id(builderStructName).Struct(
						Qual(PkgSolanaGo, "AccountMetaSlice").Tag(map[string]string{
							"bin": "-",
						}),
					)

					// func that returns a new builder for this account group:
					code.Line().Line().Func().Id("New" + builderStructName).Params().Op("*").Id(builderStructName).
						BlockFunc(func(gr *Group) {
							gr.Return().Op("&").Id(builderStructName).Block(
								Id("AccountMetaSlice").Op(":").Make(Qual(PkgSolanaGo, "AccountMetaSlice"), Lit(parentGroup.Accounts.NumAccounts())).Op(","),
							)
						}).Line().Line()

					// Method on intruction builder that accepts the accounts group builder, and copies the accounts:
					code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("Set" + ToCamel(parentGroup.Name) + "AccountsFromBuilder").
						Params(
							ListFunc(func(st *Group) {
								// Parameters:
								st.Id(ToLowerCamel(builderStructName)).Op("*").Id(builderStructName)
							}),
						).
						Params(
							ListFunc(func(st *Group) {
								// Results:
								st.Op("*").Id(insExportedName)
							}),
						).
						BlockFunc(func(gr *Group) {
							// Body:

							tpIndex := index
							// spew.Dump(parentGroup)
							for _, subAccount := range parentGroup.Accounts {
								if subAccount.IdlAccount != nil {
									exportedAccountName := ToCamel(subAccount.IdlAccount.Name)

									def := Id("inst").Dot("AccountMetaSlice").Index(Lit(tpIndex)).
										Op("=").Id(ToLowerCamel(builderStructName)).Dot(formatAccountAccessorName("Get", exportedAccountName)).Call()

									gr.Add(def)
								}
								tpIndex++
							}

							gr.Return().Id("inst")
						})
				}

				{
					exportedAccountName := ToCamel(account.Name)
					lowerAccountName := ToLowerCamel(account.Name)

					var receiverTypeName string
					if parentGroupPath == "" {
						receiverTypeName = insExportedName
					} else {
						receiverTypeName = builderStructName
					}

					code.Add(genAccountGettersSetters(
						receiverTypeName,
						account,
						groupMemberIndex,
						exportedAccountName,
						lowerAccountName,
						instruction.Accounts,
						addresses,
					))
					groupMemberIndex++
				}
				return true
			})

			file.Add(code.Line())
		}
		{
			// Declare `Build` method on instruction:
			code := Empty()

			code.Line().Line().Func().Params(Id("inst").Id(insExportedName)).Id("Build").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id("Instruction")
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					var typeIDCode Code

					GetConfig().TypeID.
						On(
							TypeIDNameSlice{
								TypeIDUvarint32,
							},
							func() {
								typeIDCode = Qual(PkgDfuseBinary, "TypeIDFromUvarint32").Call(Id("Instruction_" + insExportedName))
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDUint32,
							},
							func() {
								typeIDCode = Qual(PkgDfuseBinary, "TypeIDFromUint32").Call(Id("Instruction_"+insExportedName), Qual("encoding/binary", "LittleEndian"))
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDUint8,
							},
							func() {
								typeIDCode = Qual(PkgDfuseBinary, "TypeIDFromUint8").Call(Id("Instruction_" + insExportedName))
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDAnchor,
							},
							func() {
								typeIDCode = Id("Instruction_" + insExportedName)
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDNoType,
							},
							func() {
								// TODO
							},
						)

					body.Return().Op("&").Id("Instruction").Values(
						Dict{
							Id("BaseVariant"): Qual(PkgDfuseBinary, "BaseVariant").Values(
								Dict{
									Id("TypeID"): typeIDCode,
									Id("Impl"):   Id("inst"),
								},
							),
						},
					)
				})
			file.Add(code.Line())
		}
		{
			// Declare `ValidateAndBuild` method on instruction:
			code := Empty()

			code.Line().Line().
				Comment("ValidateAndBuild validates the instruction parameters and accounts;").
				Line().
				Comment("if there is a validation error, it returns the error.").
				Line().
				Comment("Otherwise, it builds and returns the instruction.").
				Line().
				Func().Params(Id("inst").Id(insExportedName)).Id("ValidateAndBuild").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id("Instruction")
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.If(
						Err().Op(":=").Id("inst").Dot("Validate").Call(),
						Err().Op("!=").Nil(),
					).Block(
						Return(Nil(), Err()),
					)

					body.Return(Id("inst").Dot("Build").Call(), Nil())
				})
			file.Add(code.Line())
		}
		{
			// Declare `Validate` method on instruction:
			code := Empty()

			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("Validate").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					if len(args) > 0 {
						body.Comment("Check whether all (required) parameters are set:")

						body.BlockFunc(func(paramVerifyBody *Group) {
							for _, arg := range args {
								exportedArgName := ToCamel(arg.Name)

								// Optional params can be empty.
								if arg.Type.IsIdlTypeOption() {
									continue
								}

								paramVerifyBody.If(Id("inst").Dot(exportedArgName).Op("==").Nil()).Block(
									Return(
										Qual("errors", "New").Call(Lit(Sf("%s parameter is not set", exportedArgName))),
									),
								)
							}
						})
						body.Line()
					}

					body.Comment("Check whether all (required) accounts are set:")
					body.BlockFunc(func(accountValidationBlock *Group) {
						instruction.Accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {
							exportedAccountName := ToCamel(filepath.Join(groupPath, ia.Name))

							if ia.Optional {
								accountValidationBlock.Line().Commentf(
									"[%v] = %s is optional",
									accountIndex,
									exportedAccountName,
								).Line()
							} else {
								accountValidationBlock.If(Id("inst").Dot("AccountMetaSlice").Index(Lit(accountIndex)).Op("==").Nil()).Block(
									Return(Qual("errors", "New").Call(Lit(Sf("accounts.%s is not set", exportedAccountName)))),
								)
							}

							return true
						})
					})

					body.Return(Nil())
				})
			file.Add(code.Line())
		}
		{
			// Declare `EncodeToTree(parent treeout.Branches)` method in instruction:
			code := Empty()

			code.Line().Line().Func().Params(Id("inst").Op("*").Id(insExportedName)).Id("EncodeToTree").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("parent").Qual(PkgTreeout, "Branches")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					body.Id("parent").Dot("Child").Call(Qual(PkgFormat, "Program").Call(Id("ProgramName"), Id("ProgramID"))).Op(".").
						Line().Comment("").Line().
						Id("ParentFunc").Parens(Func().Parens(Id("programBranch").Qual(PkgTreeout, "Branches")).BlockFunc(
						func(programBranchGroup *Group) {
							programBranchGroup.Id("programBranch").Dot("Child").Call(Qual(PkgFormat, "Instruction").Call(Lit(insExportedName))).Op(".").
								Line().Comment("").Line().
								Id("ParentFunc").Parens(Func().Parens(Id("instructionBranch").Qual(PkgTreeout, "Branches")).BlockFunc(
								func(instructionBranchGroup *Group) {

									instructionBranchGroup.Line().Comment("Parameters of the instruction:")

									instructionBranchGroup.Id("instructionBranch").Dot("Child").Call(Lit(Sf("Params[len=%v]", len(args)))).Dot("ParentFunc").Parens(Func().Parens(Id("paramsBranch").Qual(PkgTreeout, "Branches")).BlockFunc(func(paramsBranchGroup *Group) {
										longest := treeFindLongestNameFromFields(args)
										for _, arg := range args {
											exportedArgName := ToCamel(arg.Name)
											paramsBranchGroup.Id("paramsBranch").Dot("Child").
												Call(
													Qual(PkgFormat, "Param").Call(
														Lit(strings.Repeat(" ", longest-len(exportedArgName))+exportedArgName+StringIf(arg.Type.IsIdlTypeOption(), " (OPT)")),
														Add(CodeIf(!arg.Type.IsIdlTypeOption() && !isComplexEnum(arg.Type), Op("*"))).Id("inst").Dot(exportedArgName),
													),
												)
										}
									}))

									instructionBranchGroup.Line().Comment("Accounts of the instruction:")

									instructionBranchGroup.Id("instructionBranch").Dot("Child").Call(Lit(Sf("Accounts[len=%v]", instruction.Accounts.NumAccounts()))).Dot("ParentFunc").Parens(
										Func().Parens(Id("accountsBranch").Qual(PkgTreeout, "Branches")).BlockFunc(func(accountsBranchGroup *Group) {

											longest := treeFindLongestNameFromAccounts(instruction.Accounts)
											instruction.Accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {

												cleanedName := treeFormatAccountName(ia.Name)

												exportedAccountName := filepath.Join(groupPath, cleanedName)

												access := Id("accountsBranch").Dot("Child").Call(Qual(PkgFormat, "Meta").Call(Lit(strings.Repeat(" ", longest-len(exportedAccountName))+exportedAccountName), Id("inst").Dot("AccountMetaSlice").Dot("Get").Call(Lit(accountIndex))))
												accountsBranchGroup.Add(access)
												return true
											})
										}))
								}))
						}))
				})
			file.Add(code.Line())
		}

		{
			// Declare `MarshalWithEncoder(encoder *bin.Encoder) error` method on instruction:
			file.Add(
				genMarshalWithEncoder_struct(
					&idl,
					false,
					insExportedName,
					"",
					args,
					true,
				),
			)
		}

		{
			// Declare `UnmarshalWithDecoder(decoder *bin.Decoder) error` method on instruction:
			file.Add(
				genUnmarshalWithDecoder_struct(
					&idl,
					false,
					insExportedName,
					"",
					args,
					bin.TypeID{},
				))
		}

		{
			// Declare instruction initializer func:
			paramNames := []string{}
			for _, arg := range args {
				paramNames = append(paramNames, arg.Name)
			}
			code := Empty()
			name := "New" + insExportedName + "Instruction"
			code.Commentf("%s declares a new %s instruction with the provided parameters and accounts.", name, insExportedName)
			code.Line()
			code.Func().Id(name).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						{
							for argIndex, arg := range args {
								paramNames = append(paramNames, arg.Name)
								params.Add(func() Code {
									if argIndex == 0 {
										return Line().Comment("Parameters:")
									}
									return Empty()
								}()).Line().Id(arg.Name).Add(genTypeName(arg.Type))
							}
						}
						{
							instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
								// skip sysvars:
								if isSysVar(account.Name) {
									return true
								}
								var accountName string
								if parentGroupPath == "" {
									accountName = ToLowerCamel(account.Name)
								} else {
									accountName = ToLowerCamel(parentGroupPath + "/" + ToLowerCamel(account.Name))
								}

								if SliceContains(paramNames, accountName) {
									accountName = accountName + "Account"
								}

								params.Add(func() Code {
									if index == 0 {
										return Line().Comment("Accounts:").Line()
									}
									return Line()
								}()).Id(accountName).Qual(PkgSolanaGo, "PublicKey")
								return true
							})
						}
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id(insExportedName)
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					builder := body.Return().Id(formatBuilderFuncName(insExportedName)).Call()
					{
						for _, arg := range args {
							exportedArgName := ToCamel(arg.Name)
							builder.Op(".").Line().Id("Set" + exportedArgName).Call(Id(arg.Name))
						}
					}

					{
						declaredReceivers := []string{}
						instruction.Accounts.Walk("", nil, nil, func(parentGroupPath string, index int, parentGroup *IdlAccounts, account *IdlAccount) bool {
							// skip sysvars:
							if isSysVar(account.Name) {
								return true
							}
							var accountName string
							if parentGroupPath == "" {
								accountName = ToLowerCamel(account.Name)
							}
							//else {
							// TODO
							//}

							if SliceContains(paramNames, accountName) {
								accountName = accountName + "Account"
							}

							builderStructName := insExportedName + ToCamel(parentGroupPath) + "AccountsBuilder"
							hasNestedParent := parentGroupPath != ""
							isDeclaredReceiver := SliceContains(declaredReceivers, parentGroupPath)

							if hasNestedParent && !isDeclaredReceiver {
								declaredReceivers = append(declaredReceivers, parentGroupPath)
								builder.Op(".").Line().Id("Set" + ToCamel(parentGroup.Name) + "AccountsFromBuilder").Call(
									Line().Id("New" + builderStructName).Call().
										Add(
											DoGroup(func(gr *Group) {
												// Body:
												for subIndex, subAccount := range parentGroup.Accounts {
													if subAccount.IdlAccount != nil {
														exportedAccountName := ToCamel(subAccount.IdlAccount.Name)
														accountName = ToLowerCamel(parentGroupPath + "/" + ToLowerCamel(exportedAccountName))

														gr.Op(".").Add(func() Code {
															if subIndex == 0 {
																return Line().Line()
															}
															return Line()
														}()).Id(formatAccountAccessorName("Set", exportedAccountName)).Call(Id(accountName))

														if subIndex == len(parentGroup.Accounts)-1 {
															gr.Op(",").Line()
														}
													}
												}
											}),
										),
								)
							}

							if !hasNestedParent {
								builder.Op(".").Line().Id(formatAccountAccessorName("Set", ToCamel(account.Name))).Call(Id(accountName))
							}

							return true
						})
					}

					// builder.Op(".").Line().Id("Build").Call()
				})

			file.Add(code.Line())
		}
		////
		files = append(files, &FileWrapper{
			Name: strings.ToLower(insExportedName),
			File: file,
		})
	}

	// add configurable address map file
	{
		file := NewGoFile(idl.Metadata.Name, false)
		code := Empty().Var().Id("Addresses").Op("=").Map(String()).Qual(PkgSolanaGo, "PublicKey").Values(DictFunc(func(dict Dict) {
			for address := range addresses {
				dict[Lit(address)] = Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(address))
			}
		}))
		file.Add(code)
		files = append(files, &FileWrapper{
			Name: "addresses",
			File: file,
		})
	}

	// add constants file
	{
		file := NewGoFile(idl.Metadata.Name, false)
		code := Empty()
		for _, c := range idl.Constants {
			code.Line().Var().Id(fmt.Sprintf("CONST_%s", c.Name)).Op("=")
			typ := c.Type.GetString()
			switch typ {
			case "string":
				v, err := strconv.Unquote(c.Value)
				if err != nil {
					panic(fmt.Sprintf("failed to parse constant: %s", spew.Sdump(c)))
				}
				code.Lit(v)
			case "u16":
				v, err := strconv.ParseUint(c.Value, 10, 16)
				if err != nil {
					panic(fmt.Sprintf("failed to parse constant: %s", spew.Sdump(c)))
				}
				code.Lit(v)
			case "u64":
				v, err := strconv.ParseUint(c.Value, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("failed to parse constant: %s", spew.Sdump(c)))
				}
				code.Lit(v)
			case "i64":
				v, err := strconv.ParseInt(c.Value, 10, 64)
				if err != nil {
					panic(fmt.Sprintf("failed to parse constant: %s", spew.Sdump(c)))
				}
				code.Lit(v)
			case "pubkey":
				code.Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(c.Value))
			default:
				panic(fmt.Sprintf("unsupportd constant: %s", spew.Sdump(c)))
			}

		}
		file.Add(code)
		files = append(files, &FileWrapper{
			Name: "constants",
			File: file,
		})
	}

	{
		testFiles, err := genTestingFuncs(idl)
		if err != nil {
			return nil, err
		}
		files = append(files, testFiles...)
	}

	return files, nil
}

func genAccountGettersSetters(
	receiverTypeName string,
	account *IdlAccount,
	index int,
	exportedAccountName string,
	lowerAccountName string,
	accounts []IdlAccountItem,
	addresses map[string]string,
) Code {
	code := Empty()

	{
		code.Line().Line()
		name := formatAccountAccessorName("Set", exportedAccountName)
		code.Commentf("%s sets the %q account.", name, account.Name).Line()
		for _, doc := range account.Docs {
			code.Comment(doc).Line()
		}

		// Create account setters:
		code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(name).
			Params(
				ListFunc(func(params *Group) {
					// Parameters:
					params.Id(lowerAccountName).Qual(PkgSolanaGo, "PublicKey")
				}),
			).
			Params(
				ListFunc(func(results *Group) {
					// Results:
					results.Op("*").Id(receiverTypeName)
				}),
			).
			BlockFunc(func(body *Group) {
				// Body:
				def := Id("inst").Dot("AccountMetaSlice").Index(Lit(index)).
					Op("=").Qual(PkgSolanaGo, "Meta").Call(Id(lowerAccountName))
				if account.Writable {
					def.Dot("WRITE").Call()
				}
				if account.Signer {
					def.Dot("SIGNER").Call()
				}
				body.Add(def)

				body.Return().Id("inst")
			})
	}

	{ // create PDA helper
		/**
		func (inst *FooInstruction) FindUserTokenAmountAccountAddress(user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
			pda, bumpSeed, err = findUserTokenAmountAccountAddress(user, 0)
			return
		}

		func (inst *FooInstruction) FindUserTokenAmountAccountAddress(user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
			pda, bumpSeed, err = findUserTokenAmountAccountAddress(user, 0)
			return
		}

		func (inst *FooInstruction) MustFindUserTokenAmountAccountAddress(user ag_solanago.PublicKey) (pda ag_solanago.PublicKey) {
			pda, _, err := findUserTokenAmountAccountAddress(user)
			if err != nil {
				panic(err)
			}
			return
		}

		func (inst *FooInstruction) FindUserTokenAmountAccountAddress(user ag_solanago.PublicKey) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
			pda, bumpSeed, err = findUserTokenAmountAccountAddress(user, 0)
			return
		}

		func (inst *FooInstruction) MustFindUserTokenAmountAccountAddressWithBumpSeed(user ag_solanago.PublicKey, bumpSeed uint8) (pda ag_solanago.PublicKey) {
			pda, _, err := findUserTokenAmountAccountAddress(user, bumpSeed)
			if err != nil {
				panic(err)
			}
			return
		}

		func findUserTokenAmountAccountAddress(user ag_solanago.PublicKey, knownBumpSeed uint8) (pda ag_solanago.PublicKey, bumpSeed uint8, err error) {
			var seeds [][]byte
			seeds = append(seeds, []byte{1,2,3})
			seeds = append(seeds, user.Bytes())
			if knownBumpSeed != 0 {
				seeds = append(seeds, []byte{byte(bumpSeed)})
				pda, err = ag_solanago.CreateProgramAddress(seeds, ProgramID)
			} else {
				pda, bumpSeed, err = ag_solanago.FindProgramAddress(seeds, ProgramID)
			}
			return
		}
		*/
		if account.PDA != nil {
			code.Line().Line()
			accessorName := strings.TrimSuffix(formatAccountAccessorName("Find", exportedAccountName), "Account") + "Address"

			// find seeds
			seedValues := make([][]byte, len(account.PDA.Seeds))
			seedRefs := make([]string, len(account.PDA.Seeds))
			seedArgs := make([]string, len(account.PDA.Seeds))

			var seedProgramValue *[]byte
			if account.PDA.Program != nil {
				if account.PDA.Program.Value == nil {
					panic("cannot handle non-const type program value in PDA seeds")
				}
				seedProgramValue = &account.PDA.Program.Value
			}

		OUTER:
			for i, seedDef := range account.PDA.Seeds {
				if seedDef.Value != nil { // type: const
					seedValues[i] = seedDef.Value
				} else {
					for _, acc := range accounts {
						if acc.IdlAccount.Name == seedDef.Path {
							seedRefs[i] = ToLowerCamel(acc.IdlAccount.Name)
							continue OUTER
						}
						if seedDef.Kind == "arg" {
							seedArgs[i] = ToCamel(seedDef.Path)
							continue OUTER
						}
					}
					panic("cannot find related account path " + seedDef.Path)
				}
			}

			internalAccessorName := "find" + accessorName
			code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(internalAccessorName).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
							}
						}
						params.Id("knownBumpSeed").Uint8()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
						results.Id("bumpSeed").Uint8()
						results.Id("err").Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Add(Var().Id("seeds").Index().Index().Byte())

					for i, seedValue := range seedValues {
						if seedValue != nil {
							if utf8.Valid(seedValue) {
								if bytes.Equal(seedValue, make([]byte, len(seedValue))) {
									body.Comment("const: Pubkey.Default{}")
								} else {
									body.Commentf("const: %s", string(seedValue))
								}
							} else {
								body.Commentf("const (raw): %+v", seedValue)
							}
							body.Add(Id("seeds").Op("=").Append(Id("seeds"), Index().Byte().ValuesFunc(func(group *Group) {
								for _, v := range seedValue {
									group.LitByte(v)
								}
							})))
						} else {
							seedRef := seedRefs[i]
							if seedRef != "" {
								body.Commentf("path: %s", seedRef)
								body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id(seedRef).Dot("Bytes").Call()))
							}
							seedArg := seedArgs[i]
							if seedArg != "" {
								body.Commentf("arg: %s", seedArg)
								seedArgName := ToLowerCamel(seedArg + "Seed")
								body.Add(Id(seedArgName).Op(",").Id("err").Op(":=").Qual(PkgMsgpack, "Marshal").Call(Id("inst").Op(".").Id(seedArg)))
								body.If(
									Err().Op("!=").Nil(),
								).Block(
									Return(),
								)
								body.Add(Id("seeds").Op("=").Append(Id("seeds"), Id(seedArgName)))
							}
						}
					}

					body.Line()

					seedProgramRef := Id("ProgramID")
					if seedProgramValue != nil {
						seedProgramRef = Id("programID")
						//body.Add(Id("programID").Op(":=").Qual(PkgSolanaGo, "PublicKey").Call(Index().Byte().ValuesFunc(func(group *Group) {
						//	for _, v := range *seedProgramValue {
						//		group.LitByte(v)
						//	}
						//})))
						address := solana.PublicKeyFromBytes(*seedProgramValue).String()
						body.Add(Id("programID").Op(":=").Id("Addresses").Index(Lit(address)))
						addresses[address] = address
					}

					body.Line()

					body.Add(
						If(Id("knownBumpSeed").Op("!=").Lit(0)).BlockFunc(func(group *Group) {
							group.Add(Id("seeds").Op("=").Append(Id("seeds"), Index().Byte().Values(Byte().Call(Id("bumpSeed")))))
							group.Add(List(Id("pda"), Id("err")).Op("=").Add(Qual(PkgSolanaGo, "CreateProgramAddress").Call(Id("seeds"), seedProgramRef)))
						}).
							Else().BlockFunc(func(group *Group) {
							group.Add(List(Id("pda"), Id("bumpSeed"), Id("err")).Op("=").Add(Qual(PkgSolanaGo, "FindProgramAddress").Call(Id("seeds"), seedProgramRef)))
						}),
					)

					body.Return()
				})

			code.Line().Line()
			accessorName2 := accessorName + "WithBumpSeed"
			code.Commentf("%s calculates %s account address with given seeds and a known bump seed.", accessorName2, exportedAccountName).Line()
			code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(accessorName2).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
							}
						}
						params.Id("bumpSeed").Uint8()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
						results.Id("err").Error()
					}),
				).
				BlockFunc(func(body *Group) {
					body.Add(List(Id("pda"), Id("_"), Id("err")).Op("=").Id("inst").Dot(internalAccessorName).CallFunc(func(group *Group) {
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								group.Add(Id(seedRef))
							}
						}
						group.Add(Id("bumpSeed"))
					}))

					body.Return()
				})

			code.Line().Line()
			code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id("Must" + accessorName2).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
							}
						}
						params.Id("bumpSeed").Uint8()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
					}),
				).
				BlockFunc(func(body *Group) {
					body.Add(List(Id("pda"), Id("_"), Id("err")).Op(":=").Id("inst").Dot(internalAccessorName).CallFunc(func(group *Group) {
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								group.Add(Id(seedRef))
							}
						}
						group.Add(Id("bumpSeed"))
					}))

					body.Add(If(Id("err").Op("!=").Nil()).Block(Panic(Id("err"))))

					body.Return()
				})

			code.Line().Line()
			accessorName3 := accessorName
			code.Commentf("%s finds %s account address with given seeds.", accessorName3, exportedAccountName).Line()
			code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(accessorName3).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
							}
						}
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
						results.Id("bumpSeed").Uint8()
						results.Id("err").Error()
					}),
				).
				BlockFunc(func(body *Group) {
					body.Add(List(Id("pda"), Id("bumpSeed"), Id("err")).Op("=").Id("inst").Dot(internalAccessorName).CallFunc(func(group *Group) {
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								group.Add(Id(seedRef))
							}
						}
						group.Add(Lit(0))
					}))

					body.Return()
				})

			code.Line().Line()
			code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id("Must" + accessorName3).
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								params.Id(seedRef).Qual(PkgSolanaGo, "PublicKey")
							}
						}
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Id("pda").Qual(PkgSolanaGo, "PublicKey")
					}),
				).
				BlockFunc(func(body *Group) {
					body.Add(List(Id("pda"), Id("_"), Id("err")).Op(":=").Id("inst").Dot(internalAccessorName).CallFunc(func(group *Group) {
						for _, seedRef := range seedRefs {
							if seedRef != "" {
								group.Add(Id(seedRef))
							}
						}
						group.Add(Lit(0))
					}))

					body.Add(If(Id("err").Op("!=").Nil()).Block(Panic(Id("err"))))

					body.Return()
				})
		}
	}

	{ // Create account getters:
		code.Line().Line()
		name := formatAccountAccessorName("Get", exportedAccountName)
		if account.Optional {
			code.Commentf("%s gets the %q account (optional).", name, account.Name).Line()
		} else {
			code.Commentf("%s gets the %q account.", name, account.Name).Line()
		}
		for _, doc := range account.Docs {
			code.Comment(doc).Line()
		}
		code.Func().Params(Id("inst").Op("*").Id(receiverTypeName)).Id(name).
			Params(
				ListFunc(func(params *Group) {
					// Parameters:
				}),
			).
			Params(
				ListFunc(func(results *Group) {
					// Results:
					results.Op("*").Qual(PkgSolanaGo, "AccountMeta")
				}),
			).
			BlockFunc(func(body *Group) {
				// Body:
				body.Return(Id("inst").Dot("AccountMetaSlice").Dot("Get").Call(Lit(index)))
			})
	}

	return code
}

func genProgramBoilerplate(idl IDL) (*File, error) {
	file := NewGoFile(idl.Metadata.Name, true)
	for _, programDoc := range idl.Docs {
		file.HeaderComment(programDoc)
	}

	{
		// ProgramID variable:
		code := Empty()

		var programAddress string
		if idl.Metadata != nil && idl.Metadata.Address != "" {
			programAddress = idl.Metadata.Address // version < 0.29
		} else {
			programAddress = idl.Address
		}

		hasAddress := programAddress != ""
		code.Var().Id("ProgramID").Qual(PkgSolanaGo, "PublicKey").
			Add(
				func() Code {
					if hasAddress {
						return Op("=").Qual(PkgSolanaGo, "MustPublicKeyFromBase58").Call(Lit(programAddress))
					}
					return nil
				}(),
			)
		file.Add(code.Line())
	}
	{
		// `SetProgramID` func:
		code := Empty()
		code.Func().Id("SetProgramID").Params(Id("PublicKey").Qual(PkgSolanaGo, "PublicKey")).Block(
			Id("ProgramID").Op("=").Id("PublicKey"),
			Qual(PkgSolanaGo, "RegisterInstructionDecoder").Call(Id("ProgramID"), Id("registryDecodeInstruction")),
		)
		file.Add(code.Line())
	}
	{
		// ProgramName variable:
		code := Empty()
		programName := ToCamel(idl.Metadata.Name)
		code.Const().Id("ProgramName").Op("=").Lit(programName)
		file.Add(code.Line())
	}
	{
		// register decoder:
		code := Empty()
		code.Func().Id("init").Call().Block(
			// TODO: check for pointer instead.
			If(
				Op("!").Id("ProgramID").Dot("IsZero").Call(),
			).Block(
				Qual(PkgSolanaGo, "RegisterInstructionDecoder").Call(Id("ProgramID"), Id("registryDecodeInstruction")),
			),
		)
		file.Add(code.Line())
	}

	{
		// Instruction ID enum:
		GetConfig().TypeID.
			On(
				TypeIDNameSlice{
					TypeIDUvarint32,
					TypeIDUint32,
					TypeIDUint8,
				},
				func() {

					code := Empty()
					code.Const().Parens(
						DoGroup(func(gr *Group) {
							for instructionIndex, instruction := range idl.Instructions {
								insExportedName := ToCamel(instruction.Name)

								ins := Empty().Line()
								for _, doc := range instruction.Docs {
									ins.Comment(doc).Line()
								}
								ins.Id("Instruction_" + insExportedName)

								if instructionIndex == 0 {
									switch GetConfig().TypeID {
									case TypeIDUvarint32, TypeIDUint32:
										ins.Uint32().Op("=").Iota().Line()
									case TypeIDUint8:
										ins.Uint8().Op("=").Iota().Line()
									}
								}
								gr.Add(ins.Line().Line())
							}
						}),
					)
					file.Add(code.Line())

				},
			).
			On(
				TypeIDNameSlice{
					TypeIDAnchor,
				},
				func() {
					code := Empty()
					code.Var().Parens(
						DoGroup(func(gr *Group) {
							for _, instruction := range idl.Instructions {
								insExportedName := ToCamel(instruction.Name)

								ins := Empty().Line()
								for _, doc := range instruction.Docs {
									ins.Comment(doc).Line()
								}
								toBeHashed := sighash.ToSnakeForSighash(instruction.Name)
								if GetConfig().Debug {
									ins.Comment(Sf(`hash("%s:%s")`, bin.SIGHASH_GLOBAL_NAMESPACE, toBeHashed)).Line()
								}
								ins.Id("Instruction_" + insExportedName)

								ins.Op("=").Qual(PkgDfuseBinary, "TypeID").Call(
									Index(Lit(8)).Byte().Op("{").ListFunc(func(byteGroup *Group) {
										sighash := bin.SighashTypeID(bin.SIGHASH_GLOBAL_NAMESPACE, toBeHashed)
										if instruction.Discriminator != nil {
											sighash = *instruction.Discriminator
										}
										for _, byteVal := range sighash[:] {
											byteGroup.Lit(int(byteVal))
										}
									}).Op("}"),
								)
								gr.Add(ins.Line().Line())
							}
						}),
					)
					file.Add(code.Line())
				},
			).
			On(
				TypeIDNameSlice{
					TypeIDNoType,
				},
				func() {
					// TODO
				},
			)
	}
	{
		// Declare `InstructionIDToName` function:
		GetConfig().TypeID.
			On(
				TypeIDNameSlice{
					TypeIDUvarint32,
					TypeIDUint32,
					TypeIDUint8,
				},
				func() {
					code := Empty()
					code.Comment("InstructionIDToName returns the name of the instruction given its ID.").Line()
					code.Func().Id("InstructionIDToName").
						Params(
							func() Code {
								switch GetConfig().TypeID {
								case TypeIDUvarint32, TypeIDUint32:
									return Id("id").Uint32()
								case TypeIDUint8:
									return Id("id").Uint8()
								}
								return nil
							}(),
						).
						Params(String()).
						BlockFunc(func(body *Group) {
							body.Switch(Id("id")).BlockFunc(func(switchBlock *Group) {
								for _, instruction := range idl.Instructions {
									insExportedName := ToCamel(instruction.Name)
									switchBlock.Case(Id("Instruction_" + insExportedName)).Line().Return(Lit(insExportedName))
								}
								switchBlock.Default().Line().Return(Lit(""))
							})

						})

					file.Add(code.Line())
				},
			).
			On(
				TypeIDNameSlice{
					TypeIDAnchor,
				},
				func() {
					code := Empty()
					code.Comment("InstructionIDToName returns the name of the instruction given its ID.").Line()
					code.Func().Id("InstructionIDToName").
						Params(Id("id").Qual(PkgDfuseBinary, "TypeID")).
						Params(String()).
						BlockFunc(func(body *Group) {
							body.Switch(Id("id")).BlockFunc(func(switchBlock *Group) {
								for _, instruction := range idl.Instructions {
									insExportedName := ToCamel(instruction.Name)
									switchBlock.Case(Id("Instruction_" + insExportedName)).Line().Return(Lit(insExportedName))
								}
								switchBlock.Default().Line().Return(Lit(""))
							})

						})
					file.Add(code.Line())
				},
			).
			On(
				TypeIDNameSlice{
					TypeIDNoType,
				},
				func() {
					// TODO
				},
			)
	}

	{
		{ // Base Instruction struct:
			code := Empty()
			code.Type().Id("Instruction").Struct(
				Qual(PkgDfuseBinary, "BaseVariant"),
			)
			file.Add(code.Line())
		}
		{
			// `EncodeToTree(parent treeout.Branches)` method
			code := Empty()
			code.Func().Parens(Id("inst").Op("*").Id("Instruction")).Id("EncodeToTree").
				Params(Id("parent").Qual(PkgTreeout, "Branches")).
				Params().
				BlockFunc(func(body *Group) {
					body.If(
						List(Id("enToTree"), Id("ok")).Op(":=").Id("inst").Dot("Impl").Op(".").Parens(Qual(PkgSolanaGoText, "EncodableToTree")).
							Op(";").
							Id("ok"),
					).Block(
						Id("enToTree").Dot("EncodeToTree").Call(Id("parent")),
					).Else().Block(
						Id("parent").Dot("Child").Call(Qual("github.com/davecgh/go-spew/spew", "Sdump").Call(Id("inst"))),
					)
				})
			file.Add(code.Line())
		}
		{
			// variant definitions for the decoder:
			GetConfig().TypeID.
				On(
					TypeIDNameSlice{
						TypeIDUvarint32,
						TypeIDUint32,
						TypeIDUint8,
					},
					func() {

						code := Empty()
						code.Var().Id("InstructionImplDef").Op("=").Qual(PkgDfuseBinary, "NewVariantDefinition").
							Parens(DoGroup(func(call *Group) {
								call.Line()

								switch GetConfig().TypeID {
								case TypeIDUvarint32:
									call.Qual(PkgDfuseBinary, "Uvarint32TypeIDEncoding").Op(",").Line()
								case TypeIDUint32:
									call.Qual(PkgDfuseBinary, "Uint32TypeIDEncoding").Op(",").Line()
								case TypeIDUint8:
									call.Qual(PkgDfuseBinary, "Uint8TypeIDEncoding").Op(",").Line()
								}

								call.Index().Qual(PkgDfuseBinary, "VariantType").
									BlockFunc(func(variantBlock *Group) {
										for _, instruction := range idl.Instructions {
											// NOTE: using `ToCamel` here:
											insName := ToCamel(instruction.Name)
											insExportedName := ToCamel(instruction.Name)
											variantBlock.Block(
												List(Lit(insName), Parens(Op("*").Id(insExportedName)).Parens(Nil())).Op(","),
											).Op(",")
										}
									}).Op(",").Line()
							}))
						file.Add(code.Line())

					},
				).
				On(
					TypeIDNameSlice{
						TypeIDAnchor,
					},
					func() {
						code := Empty()
						code.Var().Id("InstructionImplDef").Op("=").Qual(PkgDfuseBinary, "NewVariantDefinition").
							Parens(DoGroup(func(call *Group) {
								call.Line()
								call.Qual(PkgDfuseBinary, "AnchorTypeIDEncoding").Op(",").Line()

								call.Index().Qual(PkgDfuseBinary, "VariantType").
									BlockFunc(func(variantBlock *Group) {
										for _, instruction := range idl.Instructions {
											// NOTE: using `ToSnakeForSighash` here (necessary for sighash computing from instruction name)
											insName := sighash.ToSnakeForSighash(instruction.Name)
											insExportedName := ToCamel(instruction.Name)
											variantBlock.Block(
												List(Id("Name").Op(":").Lit(insName), Id("Type").Op(":").Parens(Op("*").Id(insExportedName)).Parens(Nil())).Op(","),
											).Op(",")
										}
									}).Op(",").Line()
							}))
						file.Add(code.Line())
					},
				).
				On(
					TypeIDNameSlice{
						TypeIDNoType,
					},
					func() {
						// TODO
					},
				)

		}
		{
			// method to return programID:
			code := Empty()
			code.Func().Parens(Id("inst").Op("*").Id("Instruction")).Id("ProgramID").Params().
				Parens(Qual(PkgSolanaGo, "PublicKey")).
				BlockFunc(func(body *Group) {
					body.Return(
						Id("ProgramID"),
					)
				})
			file.Add(code.Line())
		}
		{
			// method to return accounts:
			code := Empty()
			code.Func().Parens(Id("inst").Op("*").Id("Instruction")).Id("Accounts").Params().
				Parens(Id("out").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")).
				BlockFunc(func(body *Group) {
					body.Return(
						Id("inst").Dot("Impl").Op(".").Parens(Qual(PkgSolanaGo, "AccountsGettable")).Dot("GetAccounts").Call(),
					)
				})
			file.Add(code.Line())
		}
		{
			// `Data() ([]byte, error)` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("Data").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Index().Byte()
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Id("buf").Op(":=").New(Qual("bytes", "Buffer"))

					body.If(
						Err().Op(":=").Qual(PkgDfuseBinary, GetConfig().Encoding._NewEncoder()).Call(Id("buf")).Dot("Encode").Call(Id("inst")).
							Op(";").
							Err().Op("!=").Nil(),
					).Block(
						Return(List(Nil(), Qual("fmt", "Errorf").Call(Lit("unable to encode instruction: %w"), Err()))),
					)
					body.Return(Id("buf").Dot("Bytes").Call(), Nil())
				})
			file.Add(code.Line())
		}
		{
			// `TextEncode(encoder *text.Encoder, option *text.Option) error` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("TextEncode").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("encoder").Op("*").Qual(PkgSolanaGoText, "Encoder")
						params.Id("option").Op("*").Qual(PkgSolanaGoText, "Option")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Return(Id("encoder").Dot("Encode").Call(Id("inst").Dot("Impl"), Id("option")))
				})
			file.Add(code.Line())
		}
		{
			// `UnmarshalWithDecoder(decoder *bin.Decoder) error` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("UnmarshalWithDecoder").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("decoder").Op("*").Qual(PkgDfuseBinary, "Decoder")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.Return(Id("inst").Dot("BaseVariant").Dot("UnmarshalBinaryVariant").Call(Id("decoder"), Id("InstructionImplDef")))
				})
			file.Add(code.Line())
		}
		{
			// `MarshalWithEncoder(encoder *bin.Encoder) error ` method:
			code := Empty()
			code.Func().Params(Id("inst").Op("*").Id("Instruction")).Id("MarshalWithEncoder").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("encoder").Op("*").Qual(PkgDfuseBinary, "Encoder")
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					GetConfig().TypeID.
						On(
							TypeIDNameSlice{
								TypeIDUvarint32,
								TypeIDUint32,
								TypeIDUint8,
							},
							func() {

								switch GetConfig().TypeID {
								case TypeIDUvarint32:
									body.Err().Op(":=").Id("encoder").Dot("WriteUVarInt").Call(Id("inst").Dot("TypeID").Dot("Uvarint32").Call())
								case TypeIDUint32:
									body.Err().Op(":=").Id("encoder").Dot("WriteUint32").Call(Id("inst").Dot("TypeID").Dot("Uint32").Call(), Qual("encoding/binary", "LittleEndian"))
								case TypeIDUint8:
									body.Err().Op(":=").Id("encoder").Dot("WriteUint8").Call(Id("inst").Dot("TypeID").Dot("Uint8").Call())
								}

							},
						).
						On(
							TypeIDNameSlice{
								TypeIDAnchor,
							},
							func() {
								body.Err().Op(":=").Id("encoder").Dot("WriteBytes").Call(Id("inst").Dot("TypeID").Dot("Bytes").Call(), False())
							},
						).
						On(
							TypeIDNameSlice{
								TypeIDNoType,
							},
							func() {
								// TODO
							},
						)

					body.If(
						Err().Op("!=").Nil(),
					).Block(
						Return(Qual("fmt", "Errorf").Call(Lit("unable to write variant type: %w"), Err())),
					)
					body.Return(Id("encoder").Dot("Encode").Call(Id("inst").Dot("Impl")))
				})
			file.Add(code.Line())
		}
		{
			// `registryDecodeInstruction` func:
			code := Empty()
			code.Func().Id("registryDecodeInstruction").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("accounts").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
						params.Id("data").Index().Byte()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Interface()
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:
					body.List(Id("inst"), Err()).Op(":=").Id("decodeInstruction").Call(Id("accounts"), Id("data"))

					body.If(
						Err().Op("!=").Nil(),
					).Block(
						Return(Nil(), Err()),
					)
					body.Return(Id("inst"), Nil())
				})
			file.Add(code.Line())
		}
		{
			// `DecodeInstruction` func:
			code := Empty()
			code.Func().Id("decodeInstruction").
				Params(
					ListFunc(func(params *Group) {
						// Parameters:
						params.Id("accounts").Index().Op("*").Qual(PkgSolanaGo, "AccountMeta")
						params.Id("data").Index().Byte()
					}),
				).
				Params(
					ListFunc(func(results *Group) {
						// Results:
						results.Op("*").Id("Instruction")
						results.Error()
					}),
				).
				BlockFunc(func(body *Group) {
					// Body:

					body.Id("inst").Op(":=").New(Id("Instruction"))

					body.If(
						Err().Op(":=").Qual(PkgDfuseBinary, GetConfig().Encoding._NewDecoder()).Call(Id("data")).Dot("Decode").Call(Id("inst")).
							Op(";").
							Err().Op("!=").Nil(),
					).Block(
						Return(
							Nil(),
							Qual("fmt", "Errorf").Call(Lit("unable to decode instruction: %w"), Err()),
						),
					)

					body.If(

						List(Id("v"), Id("ok")).Op(":=").Id("inst").Dot("Impl").Op(".").Parens(Qual(PkgSolanaGo, "AccountsSettable")).
							Op(";").
							Id("ok"),
					).BlockFunc(func(gr *Group) {
						gr.Err().Op(":=").Id("v").Dot("SetAccounts").Call(Id("accounts"))
						gr.If(Err().Op("!=").Nil()).Block(
							Return(
								Nil(),
								Qual("fmt", "Errorf").Call(Lit("unable to set accounts for instruction: %w"), Err()),
							),
						)
					})

					body.Return(Id("inst"), Nil())
				})
			file.Add(code.Line())
		}
	}

	return file, nil
}

// formatAccountAccessorName formats a name for a function that
// either gets or sets an account.
// If the RemoveAccountSuffix config flag is set, and the name already
// has an "Account" suffix, then another "Account" suffix is NOT added.
// E.g. ("Set", "Foo") => "SetFooAccount"
// E.g. ("Set", "BarAccount") => "SetBarAccount"
func formatAccountAccessorName(prefix, name string) string {
	endsWithAccount := strings.HasSuffix(strings.ToLower(name), "account")
	if !conf.RemoveAccountSuffix || !endsWithAccount {
		return prefix + name + "Account"
	}
	return prefix + name
}

func treeFindLongestNameFromFields(fields []IdlField) (ln int) {
	for _, v := range fields {
		if len(v.Name) > ln {
			ln = len(v.Name)
		}
	}
	return
}

func treeFindLongestNameFromAccounts(accounts IdlAccountItemSlice) (ln int) {
	accounts.Walk("", nil, nil, func(groupPath string, accountIndex int, parentGroup *IdlAccounts, ia *IdlAccount) bool {

		cleanedName := treeFormatAccountName(ia.Name)

		exportedAccountName := filepath.Join(groupPath, cleanedName)
		if len(exportedAccountName) > ln {
			ln = len(exportedAccountName)
		}

		return true
	})
	return
}

func treeFormatAccountName(name string) string {
	cleanedName := name
	if isSysVar(name) {
		cleanedName = strings.TrimSuffix(getSysVarName(name), "PublicKey")
	}
	if len(cleanedName) > len("account") {
		if strings.HasSuffix(cleanedName, "account") {
			cleanedName = strings.TrimSuffix(cleanedName, "account")
		} else if strings.HasSuffix(cleanedName, "Account") {
			cleanedName = strings.TrimSuffix(cleanedName, "Account")
		}
	}
	return cleanedName
}
