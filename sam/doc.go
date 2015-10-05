/*
Package sam provides editing operations on a byte slice with the syntax of the sam text editor.

See http://plan9.bell-labs.com/sys/doc/sam/sam.html for the origin and an introduction.

ADDRESSES have a start and end
		start 			end
	1	start of line 1		end of line 1
	0	start of line 1		start of line 1 (empty string)
	$	end of last line	end of last line (empty string)
	.	start of dot		end of dot
	^	beginning of current line	beginning of current line (emtpy string)

ADDRESSES can be relative (leaded by +/-) and combined (concatenated)
	/reg/	start of regex after .	end of regex after dot (short for +/reg/)
	-/reg/	start of regex before .	end of regex befor dot
	0+/reg/	start of first regex	end of first regex in file
	$-/reg/	start of last regex	end of last regex in file (search backwards from the end)
	/reg/+1	start of next line after /regexp/ after .	end of the line
	-2	start second previous line	end of second previous line to dot
	+2	start of 5th line following	end of 5th line following dot
	+-	start of line containing .	end of line containing . (short for +1-1)

ADDRESSES with comma extend from start of the first address
to the end of the last address
		start			end
	1,3	beginning of line 1 	end of line 3
	.,$	beginning of dot	end of file
	,	beginning of file	end of file (short for 0,$)

DOT
	is the current selection

COMMANDS
	i/txt/	insert txt before dot
	a/txt/	append txt after dot
	c/txt/	change dot to txt
	d	delete dot
	s/reg/txt/	substitue

LOOPS
	x/reg/ CMD 	loop through current selection, set dot and execute cmd
	0,$x/reg/ CMD	loop though whole file over, set dot to /reg/ and execute cmd
	x/reg/c/txt/	change all /reg/ to txt in the current selection
	,x/reg/c/txt/	change all /reg/ with txt in the whole file
	x/reg/ /reg2/d	delete /reg2/ following each /reg/

	x/^/ a/#/	comment selected lines with #
	x/^#/d		delete leading # for all selected lines

CONDITIONALS
	g/reg/ CMD	execute CMD if dot matches /reg/ without changing dot
	v/reg/ CMD	execute CMD if dot does not match /reg/ without changing dot
*/
package sam
