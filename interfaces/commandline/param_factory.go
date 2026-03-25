package commandline

import (
	"errors"

	"theapp/interfaces/commandline/parser"
)

// ParamFactory は、コマンドライン引数を解析してリクエストパラメータを生成する責任を持つインターフェース
type ParamFactory interface {
	Create(rawArgs [][]string) ([]interface{}, error)
}

// paramFactoryImpl は ParamFactory インターフェースの実装
type paramFactoryImpl struct {
	argumentSeparator ArgumentSeparator
	parsers           []parser.CommandArgumentParser
}

// NewParamFactory ...
func NewParamFactory() ParamFactory {
	factory := &paramFactoryImpl{
		argumentSeparator: NewArgumentSeparator(),
		parsers:           make([]parser.CommandArgumentParser, 0),
	}

	// パーサーを登録
	factory.registerParser(parser.NewOrderArgumentParser())
	factory.registerParser(parser.NewCancelArgumentParser())
	factory.registerParser(parser.NewShipArgumentParser())
	factory.registerParser(parser.NewChangeArgumentParser())
	factory.registerParser(parser.NewExpandArgumentParser())
	factory.registerParser(parser.NewInitDataArgumentParser())

	return factory
}

// registerParser パーサーを登録
func (p *paramFactoryImpl) registerParser(h parser.CommandArgumentParser) {
	p.parsers = append(p.parsers, h)
}

// Create 引数をパースしてリクエストパラメータリストを返す
func (p *paramFactoryImpl) Create(rawArgs [][]string) ([]interface{}, error) {
	// 1. 引数をコマンド別に分離
	separatedCommands, err := p.argumentSeparator.Separate(rawArgs)
	if err != nil {
		return nil, err
	}

	// 2. 各コマンドを適切なハンドラーで処理
	var paramList []interface{}
	for _, separatedCommand := range separatedCommands {
		param, err := p.parseArgument(separatedCommand)
		if err != nil {
			return nil, err
		}
		paramList = append(paramList, param)
	}

	return paramList, nil
}

// parseArgument 個別のコマンドをパース
func (p *paramFactoryImpl) parseArgument(separatedArgument SeparatedArgument) (interface{}, error) {
	// 適切なハンドラーを見つける
	for _, par := range p.parsers {
		if par.CanHandle(separatedArgument.CommandName) {
			return par.Parse(separatedArgument.Args)
		}
	}

	return nil, errors.New("no parser found for command: " + string(separatedArgument.CommandName))
}
