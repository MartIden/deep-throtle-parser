package command

import (
        "errors"
        "strconv"
)

type CommandArgs struct {
        Deep uint8
        Url  string
}

func ProcessArgs(args []string) (*CommandArgs, error) {
        if len(args) != 4 {
                return nil, errors.New("Unprocessed args")
        }

        result := &CommandArgs{}

        for i := 1; i < len(args); i++ {
                if args[i-1] == "-d" {
                        deep, err := strconv.Atoi(args[i])
                        if err != nil || deep < 1 {
                                return nil, errors.New("Deep must be positive number")
                        }
                        result.Deep = uint8(deep)
                        continue
                } else if args[i] != "-d" {
                        result.Url = args[i]
                }
        }
        return result, nil
}