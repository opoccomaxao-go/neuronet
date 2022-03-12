package core

type Neuron struct {
	Weight IOVector
	Bias   NetDataType
}

type NeuronState struct {
	ID      int         `json:"id"`
	Weights IOVector    `json:"weights"`
	Bias    NetDataType `json:"bias"`
}

func CreateNeuron(size int) *Neuron {
	return &Neuron{
		Weight: CreateIOVectorByLength(size),
	}
}

func (n *Neuron) Export(id int) NeuronState {
	return NeuronState{
		ID:      id,
		Weights: n.Weight,
		Bias:    n.Bias,
	}
}

func ImportNeuron(state NeuronState) *Neuron {
	return &Neuron{Weight: state.Weights, Bias: state.Bias}
}