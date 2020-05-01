package commands

var (
//	flagContext = map[string]map[string]string{
//	"work": {
//		flags.GlobalVarPath:         "test1",
//		flags.GlobalPipelineProfile: "test2",
//		flags.BasePath:              "test3",
//		flags.S3StateBucket:         "test4",
//	},
//}
//	config = fmt.Sprintf("/home/ksusa/.%+v.yaml", flags.ConfigFileName)
)

//func init() {
//	flags.LoadFlags()
//
//	LoadCommands()
//
//	yamlConfig, _ := yaml.Marshal(flagContext)
//
//	ioutil.WriteFile(config, yamlConfig, 0777)
//}
//
//func TestLoadFromConfig(t *testing.T) {
//	set := flag.NewFlagSet("", flag.ExitOnError)
//	set.String("config", config, "")
//
//	for profile, fls := range flagContext {
//		set.String(flags.WorkProfile, profile, "")
//		for n := range fls {
//			set.String(n, "", "")
//		}
//	}
//
//	ctx := cli.NewContext(&cli.App{Flags: flags.Flags}, set, nil)
//	if err := loadFromConfig(ctx); err != nil {
//		t.Errorf("config loading failed: %+v", err)
//		return
//	}
//
//	for n, v := range flagContext["work"] {
//		//if !c.IsSet(f) {
//		//	return fmt.Errorf("%s flag is required", f)
//		//}
//
//		val := ctx.String(n)
//		if val != v {
//
//		}
//
//		t.Log(val, v)
//	}
//}
