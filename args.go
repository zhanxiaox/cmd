package cmd

import (
	"errors"
	"os"
	"strconv"
)

type osArgs struct {
	Path    string
	Command string
	Flags   map[string]string
}

func getArgs() osArgs {
	args := osArgs{Flags: map[string]string{}}
	for k, v := range os.Args {
		if k == 0 {
			args.Path = v
		}
		if k == 1 {
			args.Command = v
		}
		if k > 1 && k%2 == 0 {
			args.Flags[v] = ""
			if len(os.Args) > k+1 {
				args.Flags[v] = os.Args[k+1]
			}
		}
	}
	return args
}

func (this *Command) mustGet(k string) (string, error) {
	v, ok := this.Flags[k]
	if ok {
		return v.value, nil
	}
	return "", errors.New("must get fail: " + k + " not found")
}

func (this *Command) shouldGet(k string) string {
	v, _ := this.Flags[k]
	return v.value
}

func (this *Command) MustGetFlagInt64(k string) (int64, error) {
	str, err := this.mustGet(k)
	if err != nil {
		return 0, errors.New("must get int64 fail: " + err.Error())
	}
	i64, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, errors.New("must get int64 fail: when parse int64 value = " + str + ",error:" + err.Error())
	}
	return i64, nil
}

func (this *Command) MustGetFlagString(k string) (string, error) {
	return this.mustGet(k)
}

func (this *Command) MustGetFlagBool(k string) (bool, error) {
	str, err := this.mustGet(k)
	if err != nil {
		return false, errors.New("must get bool fail: " + err.Error())
	}
	b, err := strconv.ParseBool(str)
	if err != nil {
		return false, errors.New("must get bool fail: when parse bool value = " + str + ",error:" + err.Error())
	}
	return b, err
}

func (this *Command) ShouldGetFlagInt64(k string) int64 {
	str := this.shouldGet(k)
	i64, _ := strconv.ParseInt(str, 10, 64)
	return i64
}

func (this *Command) ShouldGetFlagString(k string) string {
	return this.shouldGet(k)
}

func (this *Command) ShouldGetFlagBool(k string) bool {
	str := this.shouldGet(k)
	b, _ := strconv.ParseBool(str)
	return b
}
