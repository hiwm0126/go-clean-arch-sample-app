package commandline

import "errors"

// ParamFactory は、コマンドライン引数を解析してリクエストパラメータを生成する責任を持つインターフェース
type ParamFactory interface {
	Create(rawArgs [][]string) ([]interface{}, error)
}

// paramFactoryImpl は ParamFactory インターフェースの実装
type paramFactoryImpl struct {
	argumentSeparator ArgumentSeparator
	dataConverter     DataConverter
	parsers           []CommandArgumentParser
}

// NewParamFactory コマンドパーサーのコンストラクタ
func NewParamFactory() ParamFactory {
	factory := &paramFactoryImpl{
		argumentSeparator: NewArgumentSeparator(),
		dataConverter:     NewDataConverter(),
		parsers:           make([]CommandArgumentParser, 0),
	}

	// デフォルトハンドラーを登録
	factory.registerParser(NewOrderArgumentParser())
	factory.registerParser(NewCancelArgumentParser())
	factory.registerParser(NewShipArgumentParser())
	factory.registerParser(NewChangeArgumentParser())
	factory.registerParser(NewExpandArgumentParser())
	factory.registerParser(NewInitDataArgumentParser())

	return factory
}

// RegisterHandler パーサーを登録
func (p *paramFactoryImpl) registerParser(handler CommandArgumentParser) {
	p.parsers = append(p.parsers, handler)
}

// Parse 引数をパースしてリクエストパラメータリストを返す
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
	for _, parser := range p.parsers {
		if parser.CanHandle(separatedArgument.CommandName) {
			return parser.Parse(separatedArgument.Args, p.dataConverter)
		}
	}

	return nil, errors.New("no parser found for command: " + string(separatedArgument.CommandName))
}
