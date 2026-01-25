// Package wordprocessing provides APIs for creating and manipulating Word documents.
//
// Comparison and Merge:
//
// The package includes DocumentComparator and DocumentMerger for comparing and merging
// entire Word documents.
//
//	doc1, _ := wordprocessing.Open("base.docx", false)
//	doc2, _ := wordprocessing.Open("modified.docx", false)
//
//	comparator := wordprocessing.NewDocumentComparator()
//	diffs, _ := comparator.Compare(doc1, doc2)
//
//	merger := wordprocessing.NewDocumentMerger()
//	merger.Merge(doc1, doc2) // Merges doc2 into doc1
package wordprocessing
