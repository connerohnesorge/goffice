#!/bin/bash

# List of proposals in logical implementation order based on Phases
PROPOSALS=(
    # Phase 1: Foundation
    "add-validation-and-compliance"
    "add-comprehensive-element-type-coverage"
    "add-openxml-equality-comparison"
    "add-style-inheritance-and-defaults"
    "add-validation-semantic-constraints"

    # Phase 2: APIs
    "add-openxml-linq-support"
    "add-document-builder-api"
    "add-streaming-api-support"
    "add-openxml-part-reader-improvements"

    # Phase 3: Features - Word Processing
    "add-wordprocessing-advanced-features"
    "add-complete-word-api-coverage"
    "add-word-comments-enhancements"
    "add-word-content-controls"
    "add-word-custom-xml-support"
    "add-change-tracking-implementation"
    "add-hyperlink-cross-reference-system"
    "add-paragraph-id-features"
    "add-document-information-panel-support"
    "add-word-builder-api"

    # Phase 3: Features - Spreadsheet
    "add-spreadsheet-advanced-features"
    "add-full-spreadsheet-api-parity"
    "add-spreadsheet-conditional-formatting"
    "add-spreadsheet-data-validation"
    "add-spreadsheet-formula-evaluation"
    "add-spreadsheet-pivot-table-support"
    "add-workbook-calculations"
    "add-chart-advanced-features"
    "add-office-themes-support"

    # Phase 3: Features - Presentation
    "add-presentation-advanced-features"
    "add-presentation-master-support"
    "add-presentation-animation-support"
    "add-pptxgenjs-parity"

    # Phase 3: Features - DrawingML & Shapes
    "enhance-drawingml-coverage"
    "add-drawing-advanced"
    "add-strict-namespace-support"

    # Phase 3: Features - PDF Rendering
    "enhance-pdf-rendering-fidelity"
    "add-pdf-rendering-enhancements"

    # Phase 3: Features - OPC & Packaging
    "add-opc-packaging-advanced-features"
    "add-flatopc-support"
    "add-package-clone-support"
    "add-advanced-relationship-management"

    # Phase 3: Features - Framework & Architecture
    "refactor-linq-query-api"
    "add-element-event-system"

    # Phase 3: Features - Security & Others
    "add-encryption-and-protection"
    "add-macro-enabled-document-support"
    "add-metadata-properties-support"
    "add-font-management-system"
    "add-image-handling-optimization"
    "add-random-number-generator-feature"

    # Phase 4: Quality & Testing
    "expand-test-coverage"
    "add-comprehensive-test-suite"
    "add-full-validation-framework"

    # Advanced Features
    "add-mail-merge-implementation"
    "add-document-comparison-merge"
    "add-enterprise-features-parity"
)

# Loop through each proposal
for proposal in "${PROPOSALS[@]}"; do
    echo "================================================================================"
    echo "Starting Proposal: $proposal"
    echo "================================================================================"

    # Run the klaude command 3 times for each proposal
    for i in {1..3}; do
        echo "Attempt $i of 3 for $proposal..."
        
        klaude -p "/spectr:apply $proposal ensure you check if there is an existing tasks.jsonc file in spectr/changes/$proposal/ and use it if it exists to help you understand the implementation progress and what tasks remain. If it doesn't exist, create it based on the tasks.md file."
        git add . && git commit -m "Apply $i of $proposal" && git push
        
        # Check exit status if needed, though the prompt suggests running it 3 times regardless
        if [ $? -ne 0 ]; then
            echo "Warning: klaude command returned non-zero exit status on attempt $i for $proposal"
        fi
        
        echo "--------------------------------------------------------------------------------"
    done

    echo "Finished all attempts for $proposal."
    echo ""
done

echo "Completed all change proposals."
