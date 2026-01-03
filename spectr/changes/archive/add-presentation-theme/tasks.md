## 1. Implementation
- [ ] 1.1 Modify `presentation/document.go` to add a default `ThemePart` to the `PresentationPart` during initialization.
- [ ] 1.2 Update `presentation/parts/slide_master_part.go` to ensure that when a `SlideMaster` is created, it is linked to a theme. This might involve creating a relationship to the theme part.
- [ ] 1.3 Verify if `presentation.xml` needs a direct relationship to the theme (it usually doesn't, masters do).
- [ ] 1.4 Add a test case in `presentation/document_test.go` verifying the existence of `ppt/theme/theme1.xml` in a new presentation.
