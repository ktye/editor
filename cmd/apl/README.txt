# apl
--
Apl is an editor command which interpretes the selection as APL.

Lines that start with a TAB are repeated and interpreted as APL. Lines that do
not start with a TAB are ignored.

Apl interpretes the current selection, or the complete file if nothing is
selected.

Example:

    	f←{(2=+⌿0=X∘.|X)⌿X←⍳⍵}
    	f 42
    2 3 5 7 11 13 17 19 23 29 31 37 41
