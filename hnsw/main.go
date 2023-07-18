package hnsw

import (
	"fmt"
)

func main() {
	hnsw := HNSW{
		index:          []*Graph{},
		max_levels:     5,
		mult_factor:    0.62,
		efConstruction: 10,
		max_neighbors:  3,
	}
	for i := 0; i < hnsw.max_levels; i++ {
		graph := Graph{
			vertices: map[int]*Node{},
		}
		hnsw.index = append(hnsw.index, &graph)
	}

	hnsw.insert([]float64{0.1, 0.1, 0.1})
	hnsw.insert([]float64{0.2, 0.2, 0.2})
	hnsw.insert([]float64{0.3, 0.3, 0.3})
	hnsw.insert([]float64{0.4, 0.4, 0.4})

	query := []float64{0.2, 0.3, 0.2}

	nn_vals := search(hnsw.index, query, hnsw.efConstruction)
	// fmt.Println(nn_vals[0].V2)

	fmt.Println("Closest vector:")
	fmt.Println(hnsw.index[len(hnsw.index)-1].vertices[nn_vals[0].V2].vector)

	// vec_arr, question_arr := weaviate_run()
	// hnsw.create(vec_arr)

	// // ORIGINAL IDENTICAL ONE IS IN INPUT_QUERY.TXT
	// query := []float64{-1, -0.381333, -0.361916, -0.029922, -0.083385, -0.037121, -0.140229, 0.181405, 0.024734, -0.213907, 0.186477, -0.367748, -0.191971, 0.214631, -0.103155,
	// 	0.301214, -0.803744, -0.227718, -0.203787, 0.264725, 0.058073, 0.267490, 0.046337, 0.107279, -0.237854, 0.196235, 0.102541, 0.107540, -0.069553, 0.009547, 0.249763, -0.000936, 0.002275,
	// 	0.295055, 0.419565, 0.110821, 0.113490, -0.046142, 0.296329, 0.122865, -0.126267, 0.023095, -0.040004, 0.185494, -0.054055, 0.185280, 0.086308, -0.100381, -0.122253, -0.161846,
	// 	-0.076800, 0.14, -0.420127, -0.139645, -0.086043, 0.229788, 0.013104, -0.413649, -0.048755, -0.370687, 0.089419, 0.203847, 0.108975, -0.007576, 0.593015, 0.070395, -0.098459,
	// 	0.105413, -0.160375, -0.100203, 0.313930, 0.185637, -0.409707, 0.049791, 0.178786, 0.087638, -0.117298, -0.088683, -0.214002, 0.031150, -0.330146, -0.354689, 0.100124, 0.078288,
	// 	-0.217324, 0.329728, 0.198997, -0.303304, -0.309282, 0.043681, 0.100518, -0.112713, -0.021707, 0.364386, -0.112185, 0.395709, -0.083175, -0.042462, 0.533399, -0.087615, 0.116130,
	// 	0.248034, -0.219885, -0.313783, -0.003234, -0.156968, 0.192758, 0.347251, 0.167652, 0.092437, -0.077038, -0.395996, 0.009085, 0.047329, 0.165875, -0.418245, -0.186495, 0.062163,
	// 	-0.137419, -0.058573, -0.100034, -0.011791, -0.016822, -0.275621, 0.012074, -0.059855, -0.112811, 0.130563, 0.169645, -0.073759, 0.201947, 0.410747, -0.020641, 0.136406, -0.173726,
	// 	0.097144, -0.395692, -0.028602, -0.048172, -0.202710, -0.183686, -0.407371, -0.274863, -0.251035, -0.055622, -0.182175, -0.140833, 0.152275, -0.062644, 0.317231, 0.001236, -0.035688,
	// 	-0.191863, 0.247092, -0.149922, 0.131862, 0.094205, 0.225144, 0.025703, 0.290169, -0.384575, 0.204239, 0.297797, 0.396291, 0.200267, -0.339569, -0.136940, -0.052225, 0.411242,
	// 	-0.117823, 0.326824, -0.209104, 0.118624, -0.076384, -0.079706, -0.313160, -0.165852, -0.014987, 0.234074, 0.160892, 0.161103, -0.061424, 0.002757, -0.060497, 0.014424,
	// 	-0.175044, 0.008934, -0.011682, -0.018090, 0.311414, 0.168276, -0.126478, 0.197775, -0.101334, -0.108351, -0.429086, 0.293356, 0.402898, 0.010414, 0.022375, -0.181296,
	// 	0.089785, -0.141616, 0.016010, -0.048485, -0.182362, 0.191408, 0.222357, -0.378808, 0.115822, -0.130998, -0.374012, -0.370321, -0.211554, 0.004184, -0.020532, -0.165216, -0.158359,
	// 	0.083203, -0.282078, -0.012651, -0.244921, -0.066863, -0.140045, 0.141699, -0.340242, -0.053351, 0.039259, -0.044437, 0.055630, -0.059217, 0.163763, 0.213928, 0.205120, 0.157777, -0.058341,
	// 	0.090884, -0.201615, 0.123751, -0.111, -0.167558, 0.059705, 0.176185, -0.013252, 0.044728, 0.357499, -0.049555, 0.289714, -0.069825, -0.190299, -0.230479, -0.452109, 0.683444,
	// 	0.131439, 0.067736, -0.115377, 0.252148, -0.137758, -0.386849, -0.003812, 0.410928, 0.092655, 0.377023, -0.436137, 0.079425, 0.322454, -0.184361, -0.026871, 0.080180, 0.056107, 0.353761,
	// 	0.155446, 0.100516, 0.241373, 0.578041, 0.102047, 0.043927, -0.455132, -0.084355, 0.378010, -0.007609, 0.256133, 0.074233, 0.240570, -0.025764, -0.251335, -0.039906, -0.158797, -0.146657,
	// 	0.108021, -0.079393, 0.283970, -0.150876, -0.419848, 0.240125, 0.163552, 0.326806, 0.384227, -0.111319, -0.151249, -0.361685, -0.161696, 0.086195, -0.011508, 0.243702, -0.072405,
	// 	-0.031535, 0.311056, 0.231950, -0.327643, -0.154468, -0.200205, -0.246614, -0.136036, 0.099852, -0.282971, 0.123750, 0.335220, 0.018285, 0.013621, 0.157052, -0.142462, 0.051056, 0.160473, 0.201003,
	// 	0.882178, 0.046634, -0.213634, 0.105142, 0.445960, -0.144014, 0.346469, -0.252707, 0.259790, -0.071689, 0.275594, -0.490680, -0.241930, -0.117029, 0.088917, 0.003297,
	// 	-0.247074, 0.160379, -0.071612, 0.068479, -0.062621, -0.036465, 0.427177, -0.244174, -0.138298, 0.016127, 0.164683, 0.037570, 0.133779, 0.077783, -0.103568, -0.237310, 0.077587,
	// 	0.077033, -0.094969, -0.112353, -0.147646, -0.305772, 0.301303, 0.019364, 0.278044, -0.274390, 0.098462, -0.086774, 0.057400, -0.060618, -0.012532, 0.268643, -0.241573, -0.094204, 0.074196,
	// 	-0.044285, 0.322739, -0.047965, -0.205142, 0.263196, 0.187069, 0.802158, -0.330193,
	// }

	// nn_vals := search(hnsw.index, query, hnsw.efConstruction)
	// // fmt.Println(nn_vals[0].V2)

	// fmt.Println("Closest vector:")
	// fmt.Println(hnsw.index[len(hnsw.index)-1].vertices[nn_vals[0].V2].vector)
	// fmt.Println(question_arr[position(vec_arr, hnsw.index[len(hnsw.index)-1].vertices[nn_vals[0].V2].vector)])

}
