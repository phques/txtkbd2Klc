// txtkbd2Klc project  
// Copyright 2016 Philippe Quesnel  
// Licensed under the Academic Free License version 3.0 


Quick & dirty Go program to read a text file definition of a keyboard layout 
and output in a format useable with Microsoft keyboard layout editor file format

Takes something like this from a text file:
(only the 30 main keys, whitespace will be removed)

		B L O U / J D C P Y
		H R E A ? G T S N I
		K X < > Z W M F V Q

		b l o u ; j d c p y  
		h r e a , g t s n i  
		k x : . z w m f v q  

and outputs this:

		//  bB lL oO uU ;/ jJ dD cC pP yY  
		//  hH rR eE aA ,? gG tT sS nN iI  
		//  kK xX :< .> zZ wW mM fF vV qQ  

		10      B               1       b       B               // b B
		11      L               1       l       L               // l L
		12      O               1       o       O               // o O
		13      U               1       u       U               // u U
		14      OEM_1           0       003b    002f            // ; /
		15      J               1       j       J               // j J
		16      D               1       d       D               // d D
		17      C               1       c       C               // c C
		18      P               1       p       P               // p P
		19      Y               1       y       Y               // y Y
		
		1e      H               1       h       H               // h H
		1f      R               1       r       R               // r R
		20      E               1       e       E               // e E
		21      A               1       a       A               // a A
		22      OEM_COMMA       0       002c    003f            // , ?
		23      G               1       g       G               // g G
		24      T               1       t       T               // t T
		25      S               1       s       S               // s S
		26      N               1       n       N               // n N
		27      I               1       i       I               // i I
		
		2c      K               1       k       K               // k K
		2d      X               1       x       X               // x X
		// this VK is already used ! manually fix it
		2e      OEM_1           0       003a    003c            // : <
		2f      OEM_PERIOD      0       002e    003e            // . >
		30      Z               1       z       Z               // z Z
		31      W               1       w       W               // w W
		32      M               1       m       M               // m M
		33      F               1       f       F               // f F
		34      V               1       v       V               // v V
		35      Q               1       q       Q               // q Q

Just copy / paste the layout entries into a .KLC file, replacing the required keys.  
(see templateKLC.klc, the 30 main keys are already removed)

### Note
Since VK_xx codes are *per key*, if the symbols from a QWERTY key are placed on two different keys (eg. /?), then one of the keys will need to use a different VK code, otherwise Microsoft layout editor will have problems (with duplicate VK_ entries).

This needs to be handled manually, this program does not handle it.
It outputs unused VK_ keys to help in the process though

