const FORTRAN = require("tree-sitter-fortran/grammar")

const PREC = Object.assign(FORTRAN.PREC, {})

module.exports = grammar(FORTRAN, {
  name: 'fixed_form_fortran',

  externals: ($, original) => original.concat([
    $._comment_character
  ]),

  rules: {
    comment: $ => seq($._comment_character, token(/.*/)),
  }
})
