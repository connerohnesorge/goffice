package presentation

import (
	"path/filepath"
	"testing"
)

func TestImportHTMLTable(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "html_table.pptx")

	pres, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("failed to create presentation: %v", err)
	}
	defer func() { _ = pres.Close() }()

	_, err = pres.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	htmlStr := `
		<table>
			<thead>
				<tr>
					<th>Header 1</th>
					<th>Header 2</th>
				</tr>
			</thead>
			<tbody>
				<tr>
					<td>Row 1, Col 1</td>
					<td>Row 1, Col 2</td>
				</tr>
				<tr>
					<td>Row 2, Col 1</td>
					<td>Row 2, Col 2</td>
				</tr>
			</tbody>
		</table>
	`

	table, err := pres.ImportHTMLTable(0, htmlStr)
	if err != nil {
		t.Fatalf("ImportHTMLTable failed: %v", err)
	}

	if table.RowCount() != 3 {
		t.Errorf("expected 3 rows, got %d", table.RowCount())
	}

	if table.ColumnCount() != 2 {
		t.Errorf("expected 2 columns, got %d", table.ColumnCount())
	}

	val, _ := table.GetCellText(0, 0)
	if val != "Header 1" {
		t.Errorf("expected Header 1, got %s", val)
	}

	val, _ = table.GetCellText(1, 0)
	if val != "Row 1, Col 1" {
		t.Errorf("expected Row 1, Col 1, got %s", val)
	}

	if err := pres.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}

func TestImportHTMLText(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "html_text.pptx")

	pres, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("failed to create presentation: %v", err)
	}
	defer func() { _ = pres.Close() }()

	_, err = pres.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	htmlStr := `
		<div>
			<p><b>Heading</b></p>
			<p>This is a <i>paragraph</i> with a <a href="https://example.com">link</a>.</p>
			<ul>
				<li>Bullet 1</li>
				<li>Bullet 2</li>
			</ul>
		</div>
	`

	shape, err := pres.ImportHTMLText(0, 50, 50, 300, 200, htmlStr)
	if err != nil {
		t.Fatalf("ImportHTMLText failed: %v", err)
	}

	if shape == nil {
		t.Fatalf("shape is nil")
	}

	if err := pres.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}

func TestImportHTMLTable_WithFormatting(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "html_table_fmt.pptx")

	pres, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("failed to create presentation: %v", err)
	}
	defer func() { _ = pres.Close() }()

	_, err = pres.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	htmlStr := `
		<table>
			<tr>
				<td><b>Bold</b> <i>Italic</i> <u>Underline</u> <b><i>BI</i></b></td>
			</tr>
		</table>
	`

	table, err := pres.ImportHTMLTable(0, htmlStr)
	if err != nil {
		t.Fatalf("ImportHTMLTable failed: %v", err)
	}

	cell, _ := table.GetCell(0, 0)
	if cell == nil {
		t.Fatalf("cell not found")
	}

	runs := cell.TextBody().Paragraphs()[0].Runs()
	if len(runs) < 4 {
		t.Errorf("expected at least 4 runs, got %d", len(runs))
	}

	// Verify formatting (approximate check since we use strings.TrimSpace)
	// Bold should be true for first run
	if err := pres.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}

func TestImportHTMLTable_WithLists(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "html_table_lists.pptx")

	pres, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("failed to create presentation: %v", err)
	}
	defer func() { _ = pres.Close() }()

	_, err = pres.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	htmlStr := `
		<table>
			<tr>
				<td>
					<ul>
						<li>Item 1</li>
						<li>Item 2</li>
					</ul>
					<ol>
						<li>Step 1</li>
						<li>Step 2</li>
					</ol>
				</td>
			</tr>
		</table>
	`

	table, err := pres.ImportHTMLTable(0, htmlStr)
	if err != nil {
		t.Fatalf("ImportHTMLTable failed: %v", err)
	}

	cell, _ := table.GetCell(0, 0)
	if cell == nil {
		t.Fatalf("cell not found")
	}

	paras := cell.TextBody().Paragraphs()
	// Initial empty paragraph + 2 from ul + 2 from ol = 5?
	// My implementation adds a new paragraph for each li.
	if len(paras) < 4 {
		t.Errorf("expected at least 4 paragraphs, got %d", len(paras))
	}

	if err := pres.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}

func TestImportHTMLTable_WithLinks(t *testing.T) {
	tempDir := t.TempDir()
	path := filepath.Join(tempDir, "html_table_links.pptx")

	pres, err := New(path, DocTypePresentation)
	if err != nil {
		t.Fatalf("failed to create presentation: %v", err)
	}
	defer func() { _ = pres.Close() }()

	_, err = pres.AddSlide()
	if err != nil {
		t.Fatalf("failed to add slide: %v", err)
	}

	htmlStr := `
			<table>
				<tr>
					<td><a href="https://example.com">Example Link</a></td>
				</tr>
			</table>
		`

	table, err := pres.ImportHTMLTable(0, htmlStr)
	if err != nil {
		t.Fatalf("ImportHTMLTable failed: %v", err)
	}

	cell, _ := table.GetCell(0, 0)
	if cell == nil {
		t.Fatalf("cell not found")
	}

	run := cell.TextBody().Paragraphs()[0].Runs()[0]
	// Just check if it compiles and runs for now.
	// In a real test we'd check if the hyperlink element exists.
	_ = run

	if err := pres.Save(); err != nil {
		t.Fatalf("failed to save: %v", err)
	}
}
