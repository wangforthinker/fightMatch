package match


const(
	NoDiff = 0
)

//匹配时,计算差异性,并给出评分

type Difference interface {
	//compute the diff between scorers, return the avg score and diff value
	//diff value in[0,100], 0 means no diff, 100 means most diff
 	Diff([]Scorer) (avg int,diff int)
}