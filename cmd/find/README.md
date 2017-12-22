# find
--
Find is an editor command which finds and replaces text.

If called with a single argument, it sets the selection to all occurences of the
regular expression passed as the argument.

If called without an argument, it uses the FirstSelectedText as the regular
expression.

If called with 2 arguments it uses the second argument as the replace text.
Within the replace text, $ signs are interpreted as in regexp.ReplaceAllString.

Find (and replace) is only applied to the selected text, if there is any
selection and find has not been called without arguments.
