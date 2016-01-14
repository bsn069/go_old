package bsn_input

var gInput *sInput

func instance() *sInput {
	if gInput == nil {
		gInput = &sInput{
			m_tMod2Chan:          make(tMod2Chan),
			m_cmd:                new(SCmd),
			m_quit:               false,
			m_tUpperName2RegName: make(map[string]string),
			m_strUseMod:          "",
			m_bRuning:            false,
		}
	}
	return gInput
}

func run() {
	instance().run()
}
