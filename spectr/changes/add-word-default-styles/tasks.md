## 1. Implementation
- [ ] 1.1 Update `wordprocessing/parts/styles_part.go` to include a `InitializeDefault` method that sets minimal XML content (docDefaults, Normal style).
- [ ] 1.2 Modify `wordprocessing/document.go`'s `initializeDocument` function to call `AddStylesPart()` and then `InitializeDefault()` on the new part.
- [ ] 1.3 Add a unit test in `wordprocessing/document_test.go` to verify that `New()` creates a document with a `styles.xml` part.
- [ ] 1.4 Add a roundtrip test to ensure the default styles are preserved and valid.
