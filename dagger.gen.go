// Code generated by dagger. DO NOT EDIT.

package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log/slog"
	"os"

	"github.com/goharbor/harbor-cli/internal/dagger"
	"github.com/goharbor/harbor-cli/internal/telemetry"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
	"go.opentelemetry.io/otel/trace"
)

var dag = dagger.Connect()

func Tracer() trace.Tracer {
	return otel.Tracer("dagger.io/sdk.go")
}

// used for local MarshalJSON implementations
var marshalCtx = context.Background()

// called by main()
func setMarshalContext(ctx context.Context) {
	marshalCtx = ctx
	dagger.SetMarshalContext(ctx)
}

type DaggerObject = dagger.DaggerObject

type ExecError = dagger.ExecError

// ptr returns a pointer to the given value.
func ptr[T any](v T) *T {
	return &v
}

// convertSlice converts a slice of one type to a slice of another type using a
// converter function
func convertSlice[I any, O any](in []I, f func(I) O) []O {
	out := make([]O, len(in))
	for i, v := range in {
		out[i] = f(v)
	}
	return out
}

func (r HarborCli) MarshalJSON() ([]byte, error) {
	var concrete struct{}
	return json.Marshal(&concrete)
}

func (r *HarborCli) UnmarshalJSON(bs []byte) error {
	var concrete struct{}
	err := json.Unmarshal(bs, &concrete)
	if err != nil {
		return err
	}
	return nil
}

func main() {
	ctx := context.Background()

	// Direct slog to the new stderr. This is only for dev time debugging, and
	// runtime errors/warnings.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stderr, &slog.HandlerOptions{
		Level: slog.LevelWarn,
	})))

	if err := dispatch(ctx); err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func dispatch(ctx context.Context) error {
	ctx = telemetry.InitEmbedded(ctx, resource.NewWithAttributes(
		semconv.SchemaURL,
		semconv.ServiceNameKey.String("dagger-go-sdk"),
		// TODO version?
	))
	defer telemetry.Close()

	// A lot of the "work" actually happens when we're marshalling the return
	// value, which entails getting object IDs, which happens in MarshalJSON,
	// which has no ctx argument, so we use this lovely global variable.
	setMarshalContext(ctx)

	fnCall := dag.CurrentFunctionCall()
	parentName, err := fnCall.ParentName(ctx)
	if err != nil {
		return fmt.Errorf("get parent name: %w", err)
	}
	fnName, err := fnCall.Name(ctx)
	if err != nil {
		return fmt.Errorf("get fn name: %w", err)
	}
	parentJson, err := fnCall.Parent(ctx)
	if err != nil {
		return fmt.Errorf("get fn parent: %w", err)
	}
	fnArgs, err := fnCall.InputArgs(ctx)
	if err != nil {
		return fmt.Errorf("get fn args: %w", err)
	}

	inputArgs := map[string][]byte{}
	for _, fnArg := range fnArgs {
		argName, err := fnArg.Name(ctx)
		if err != nil {
			return fmt.Errorf("get fn arg name: %w", err)
		}
		argValue, err := fnArg.Value(ctx)
		if err != nil {
			return fmt.Errorf("get fn arg value: %w", err)
		}
		inputArgs[argName] = []byte(argValue)
	}

	result, err := invoke(ctx, []byte(parentJson), parentName, fnName, inputArgs)
	if err != nil {
		return fmt.Errorf("invoke: %w", err)
	}
	resultBytes, err := json.Marshal(result)
	if err != nil {
		return fmt.Errorf("marshal: %w", err)
	}
	if err = fnCall.ReturnValue(ctx, dagger.JSON(resultBytes)); err != nil {
		return fmt.Errorf("store return value: %w", err)
	}
	return nil
}
func invoke(ctx context.Context, parentJSON []byte, parentName string, fnName string, inputArgs map[string][]byte) (_ any, err error) {
	_ = inputArgs
	switch parentName {
	case "HarborCli":
		switch fnName {
		case "Echo":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var stringArg string
			if inputArgs["stringArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["stringArg"]), &stringArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg stringArg", err))
				}
			}
			return (*HarborCli).Echo(&parent, stringArg), nil
		case "ContainerEcho":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var stringArg string
			if inputArgs["stringArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["stringArg"]), &stringArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg stringArg", err))
				}
			}
			return (*HarborCli).ContainerEcho(&parent, stringArg), nil
		case "GrepDir":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var directoryArg *dagger.Directory
			if inputArgs["directoryArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["directoryArg"]), &directoryArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg directoryArg", err))
				}
			}
			var pattern string
			if inputArgs["pattern"] != nil {
				err = json.Unmarshal([]byte(inputArgs["pattern"]), &pattern)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg pattern", err))
				}
			}
			return (*HarborCli).GrepDir(&parent, ctx, directoryArg, pattern)
		case "LintCode":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var directoryArg *dagger.Directory
			if inputArgs["directoryArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["directoryArg"]), &directoryArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg directoryArg", err))
				}
			}
			return (*HarborCli).LintCode(&parent, ctx, directoryArg), nil
		case "BuildHarbor":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var directoryArg *dagger.Directory
			if inputArgs["directoryArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["directoryArg"]), &directoryArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg directoryArg", err))
				}
			}
			return (*HarborCli).BuildHarbor(&parent, ctx, directoryArg), nil
		case "PullRequest":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var directoryArg *dagger.Directory
			if inputArgs["directoryArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["directoryArg"]), &directoryArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg directoryArg", err))
				}
			}
			var githubToken string
			if inputArgs["githubToken"] != nil {
				err = json.Unmarshal([]byte(inputArgs["githubToken"]), &githubToken)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg githubToken", err))
				}
			}
			(*HarborCli).PullRequest(&parent, ctx, directoryArg, githubToken)
			return nil, nil
		case "Release":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var directoryArg *dagger.Directory
			if inputArgs["directoryArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["directoryArg"]), &directoryArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg directoryArg", err))
				}
			}
			var githubToken string
			if inputArgs["githubToken"] != nil {
				err = json.Unmarshal([]byte(inputArgs["githubToken"]), &githubToken)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg githubToken", err))
				}
			}
			(*HarborCli).Release(&parent, ctx, directoryArg, githubToken)
			return nil, nil
		case "DockerPublish":
			var parent HarborCli
			err = json.Unmarshal(parentJSON, &parent)
			if err != nil {
				panic(fmt.Errorf("%s: %w", "failed to unmarshal parent object", err))
			}
			var directoryArg *dagger.Directory
			if inputArgs["directoryArg"] != nil {
				err = json.Unmarshal([]byte(inputArgs["directoryArg"]), &directoryArg)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg directoryArg", err))
				}
			}
			var regUsername string
			if inputArgs["regUsername"] != nil {
				err = json.Unmarshal([]byte(inputArgs["regUsername"]), &regUsername)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg regUsername", err))
				}
			}
			var regPassword *dagger.Secret
			if inputArgs["regPassword"] != nil {
				err = json.Unmarshal([]byte(inputArgs["regPassword"]), &regPassword)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg regPassword", err))
				}
			}
			var privateKey *dagger.Secret
			if inputArgs["privateKey"] != nil {
				err = json.Unmarshal([]byte(inputArgs["privateKey"]), &privateKey)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg privateKey", err))
				}
			}
			var password *dagger.Secret
			if inputArgs["password"] != nil {
				err = json.Unmarshal([]byte(inputArgs["password"]), &password)
				if err != nil {
					panic(fmt.Errorf("%s: %w", "failed to unmarshal input arg password", err))
				}
			}
			return (*HarborCli).DockerPublish(&parent, ctx, directoryArg, regUsername, regPassword, privateKey, password), nil
		default:
			return nil, fmt.Errorf("unknown function %s", fnName)
		}
	case "":
		return dag.Module().
			WithObject(
				dag.TypeDef().WithObject("HarborCli").
					WithFunction(
						dag.Function("Echo",
							dag.TypeDef().WithKind(dagger.StringKind)).
							WithArg("stringArg", dag.TypeDef().WithKind(dagger.StringKind))).
					WithFunction(
						dag.Function("ContainerEcho",
							dag.TypeDef().WithObject("Container")).
							WithDescription("Returns a container that echoes whatever string argument is provided").
							WithArg("stringArg", dag.TypeDef().WithKind(dagger.StringKind))).
					WithFunction(
						dag.Function("GrepDir",
							dag.TypeDef().WithKind(dagger.StringKind)).
							WithDescription("Returns lines that match a pattern in the files of the provided Directory").
							WithArg("directoryArg", dag.TypeDef().WithObject("Directory")).
							WithArg("pattern", dag.TypeDef().WithKind(dagger.StringKind))).
					WithFunction(
						dag.Function("LintCode",
							dag.TypeDef().WithObject("Container")).
							WithArg("directoryArg", dag.TypeDef().WithObject("Directory"))).
					WithFunction(
						dag.Function("BuildHarbor",
							dag.TypeDef().WithObject("Directory")).
							WithArg("directoryArg", dag.TypeDef().WithObject("Directory"))).
					WithFunction(
						dag.Function("PullRequest",
							dag.TypeDef().WithKind(dagger.VoidKind).WithOptional(true)).
							WithArg("directoryArg", dag.TypeDef().WithObject("Directory")).
							WithArg("githubToken", dag.TypeDef().WithKind(dagger.StringKind))).
					WithFunction(
						dag.Function("Release",
							dag.TypeDef().WithKind(dagger.VoidKind).WithOptional(true)).
							WithArg("directoryArg", dag.TypeDef().WithObject("Directory")).
							WithArg("githubToken", dag.TypeDef().WithKind(dagger.StringKind))).
					WithFunction(
						dag.Function("DockerPublish",
							dag.TypeDef().WithKind(dagger.StringKind)).
							WithArg("directoryArg", dag.TypeDef().WithObject("Directory")).
							WithArg("regUsername", dag.TypeDef().WithKind(dagger.StringKind)).
							WithArg("regPassword", dag.TypeDef().WithObject("Secret")).
							WithArg("privateKey", dag.TypeDef().WithObject("Secret")).
							WithArg("password", dag.TypeDef().WithObject("Secret")))), nil
	default:
		return nil, fmt.Errorf("unknown object %s", parentName)
	}
}
