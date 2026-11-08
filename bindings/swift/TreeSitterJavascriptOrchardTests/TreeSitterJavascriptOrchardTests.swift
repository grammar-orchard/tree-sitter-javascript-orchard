import XCTest
import SwiftTreeSitter
import TreeSitterJavascriptOrchard

final class TreeSitterJavascriptOrchardTests: XCTestCase {
    func testCanLoadGrammar() throws {
        let parser = Parser()
        let language = Language(language: tree_sitter_javascript_orchard())
        XCTAssertNoThrow(try parser.setLanguage(language),
                         "Error loading JavaScript grammar")
    }
}
