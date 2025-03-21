import XCTest
import SwiftTreeSitter
import TreeSitterFixedFormFortran

final class TreeSitterFixedFormFortranTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_fixed_form_fortran())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading FixedFormFortran grammar")
    }
}
