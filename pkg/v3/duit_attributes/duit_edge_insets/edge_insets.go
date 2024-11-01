package duit_edge_insets

type EdgeInsetsAll float32

// Assign EdgeInsets in order -> left, top, right, bottom
type EdgeInsetsLTRB [4]float32

// Assign EdgeInsets in order -> vertival, horizontal
type EdgeInsetsSymmentric [2]float32

type EdgeInsets interface {
	EdgeInsetsAll | EdgeInsetsLTRB | EdgeInsetsSymmentric
}
