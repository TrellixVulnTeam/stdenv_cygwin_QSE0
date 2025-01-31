GOOF----LE-8-2.0�&      ] t 4  h�      ] g  guile�	 �	g  define-module*�	 �	 �	g  system�	g  vm�	g  inspect�		 �	
g  filenameS�	f  system/vm/inspect.scm�	g  importsS�	g  base�	g  pmatch�	 �	 �	g  syntax�	 �	 �	 �	 �	g  frame�	 �	 �	g  language�	g  assembly�	g  disassemble�	 �	g  selectS�	g  %disassemble�	��	  �	!  �	"g  ice-9�	#g  rdelim�	$"# �	%$ �	&g  pretty-print�	'"& �	(' �	)g  format�	*") �	+* �	,g  program�	-, �	.- �	/!%(+. 	�	0g  exportsS�	1 �	2g  set-current-module�	32 �	42 �	5g  make-hash-table�	6g  hash-for-each�	7g  
hashq-set!�	8g  	hashq-ref�	9g  reverse-hashq�	:g  catch�	;g  wrong-number-of-args�	<g  keyword-argument-error�	=g  catch-bad-arguments�	>g  eof-object?�	?g  char=?�	@g  	read-char�	Ag  repl-reader�	Bg  char-whitespace?�	Cg  unread-char�	Dg  read�	Eg  reverse�	Fg  	read-args�	Gg  make-module�	Hg  throw�	Ig  quit�	Jg  set-procedure-property!�	Kg  name�	Lg  module-define!�	Mg  module-add!�	Ng  q�	Og  module-local-variable�	Pg  continue�	Qg  cont�	Rg  c�	Sg  print�	Tg  p�	Ug  write�	Vg  w�	Wg  display�	Xg  d�	Yf  Error disassembling object: ~a
�	Zg  x�	[g  module-obarray�	\g  procedure-name�	]f  ~a~{ ~:@(~a~)~}~?~%~a~&~%�	^g  program-lambda-list�	_f  "~#[~:;~40t(aliases: ~@{~a~^, ~})~]�	`g  delq�	ag  procedure-documentation�	bf  Invalid command ~s.~%�	cf  #Try `help' for a list of commands~%�	dg  sort�	eg  hash-map->list�	fg  string<?�	gg  symbol->string�	hf  Available commands:~%~%�	ig  for-each�	jg  help�	kg  h�	lg  ?�	mg  apply�	ng  current-error-port�	of  )Invalid arguments to ~a. Try `help ~a'.~%�	pg  newline�	qf  $~&Unknown command: ~a. Try `help'.~%�	rg  save-module-excursion�	sf  ~20@y inspect> �C 5      h   y   ]4	
/015 44 >  "  G   5678       h   �   ]L  4L 5�6 �       g  k
		 g  v		  g  filenamef  system/vm/inspect.scm�
	%	��		&	!��		&	2��		&	!��		&	��		&	�� 			   C      h(   �   ]	45 4O  >  "  G  C   �       g  h
		% g  ret		%  g  filenamef  system/vm/inspect.scm�
	"
��		#	��		#	��	
	$	�� 		%  g  nameg  reverse-hashq� C9R:;:<h   s   - 1 3 L 6    k       g  k
			 g  args			  g  filenamef  system/vm/inspect.scm�
	/	��		0	
�� 			
   C  h   X   ] LL O 6       P       g  filenamef  system/vm/inspect.scm�
	,	��		-	��		-	�� 		
   C      h   s   - 1 3 L 6    k       g  k
			 g  args			  g  filenamef  system/vm/inspect.scm�
	1	��		2	�� 			
   C  h   �   ] O O 6�       g  thunk
		 g  bad-args-thunk		  g  filenamef  system/vm/inspect.scm�
	*
��		+		��		+	�� 			  g  nameg  catch-bad-arguments� C=R>?@ABCDE 	 h@  �  ]!"  �45$   C4
5$  4 5"���45$  4 5"���4>  "  G  4 5 4 5"  "  �45$  64
5$  645$  4 5"���4>  "  G  4 5�4 5"��v4 5"����      g  prompt
	@ g  chr	 � g  reader		&	4 g  reader		G	U g  reader		o	} g  tok		} � g  reader	 � � g  out �* g  chr	 �* g  reader	 � � g  reader	 � g  tok	* g  reader	" g  reader,:  g  filenamef  system/vm/inspect.scm�
	4
��		9	��		;	��		:	��		;	��		<	��	$	:	��	&	8	��	)	6	��	:	<	��	;	=	��	E	:	��	G	8	��	J	6	��	[	=	��	\	?	��	o	@	��	r	6	��	}	@	�� �	A	�� �	8	�� �	6	�� �	A	�� �	B	�� �	D	�� �	C	�� �	D	�� �	E	�� �	C	�� �	E	�� �	F	�� �	C	�� �	8	�� �	6	�� �	F	�� �	H	�� �	I	�� 	6	��	I	��	J	��	8	��	6	��*	J	��,	8	��/	6	��@	K	�� /	@  g  nameg  	read-args� CFRGHI       h   �   ] 6�       g  filenamef  system/vm/inspect.scm�
	a	��		c	��		c	�� 		
  g  nameg  c�g  documentationf  Quit the inspector.� CJKILMNOPQR&       h   �   ] L 6�       g  filenamef  system/vm/inspect.scm�
	e	��		g	�� 		
  g  nameg  c�g  documentationf  .Print the current object using `pretty-print'.� CSTh   �   ] ML 6�       g  filenamef  system/vm/inspect.scm�
	i	��		k	�� 		
  g  nameg  c�g  documentationf  'Print the current object using `write'.� CUV       h   �   ] ML 6�       g  filenamef  system/vm/inspect.scm�
	m	��		o	�� 		
  g  nameg  c�g  documentationf  )Print the current object using `display'.� CWX:      h   P   ] L 6H       g  filenamef  system/vm/inspect.scm�
	t	��		u	
�� 		
   C)Y        h   m   -  1  3  6      e       g  args
			  g  filenamef  system/vm/inspect.scm�
	v	��		w	��		w	
�� 			


   C        h   �   ] L O 6�       g  filenamef  system/vm/inspect.scm�
	q	��		s	�� 		
  g  nameg  c�g  documentationf  GDisassemble the current object, which should be objcode or a procedure.� CZ9[O\)]^_`8a 
     hH      ]4L  5454544L554	56          g  cmd
		D g  v		D g  p			D g  canonical-name			D  g  filenamef  system/vm/inspect.scm�
	|	��		}	��		}	
��		~	��		}	
��			!��		}	
��	  �	��	# �	#��	+ �	��	, �	��	1 �	)��	; �	��	< �	��	D �	�� 		D  g  nameg  help-cmd� C)bcO\]^_`8ade\        h   z   ] 6       r       g  k
			 g  v			  g  filenamef  system/vm/inspect.scm�
 �	��	 �	+��		 �	�� 				   Cfg      h   �   ]4 5456    {       g  x
		 g  y		  g  filenamef  system/vm/inspect.scm�
 �	��	 �	$��	 �	$��	 �	�� 			   Chi      h�   i  -  . , 3  #   44L 55O L Q  $  ~"  4 >  "  G  6 �$  S4L  5$  B4L  545	4
54455456"���"���44554>  "  G  6    a      g  cmd
	 � g  rhash	! � g  help-cmd		, � g  v		v � g  p		{ � g  canonical-name	 � � g  names	 � �  g  filenamef  system/vm/inspect.scm�
	y	��		{	��		{	"��	!	{	��	!	{	��	9 �	��	> �	��	C �	��	J �	��	W �	��	Y �	��	Y �	
��	\ �	��	` �	
��	a �	��	m �	��	n	}	��	v	}	
��	{	~	��	{	}	
��	~		!�� �	}	
�� � �	�� � �	#�� � �	�� � �	�� � �	)�� � �	�� � �	�� � �	�� � �	�� � �	�� � �	�� � �	
�� � �	�� � �	�� � �	�� � �	�� '	 �
  g  nameg  c�g  documentationf  Show this help message.� Cjkl:mO=     h   [   ] LL @       S       g  filenamef  system/vm/inspect.scm�
 �	��	 �	��		 �	�� 			
   C)no\ h    v   ] 45 4L 54L 56     n       g  filenamef  system/vm/inspect.scm�
 �	��	 �	��	 �	��	 �	��	 �	0��	 �	�� 		
   C>pHI)nq h�   2  - 1 3  �$  4L  5"  $  O O 64 5$  4>   "  G  64	4
5  >  "  G  C     *      g  cmd
			{ g  args			{ g  t			{ g  proc		)	?  g  filenamef  system/vm/inspect.scm�
 �	��	 �	��	 �	��	 �	��	 �	��	) �	��	) �	��	? �	��	@ �	��	J �	��	K �	��	^ �	��	` �	��	a �	��	d �	��	j �	��	q �	�� 			{
  g  nameg  handle� Cr2F)s       h   k   ] L 6     c       g  filenamef  system/vm/inspect.scm�
	^	��		_	��		_	�� 		
  g  nameg  prompt� C  h    [   ] 4L>  "  G  L O 6 S       g  filenamef  system/vm/inspect.scm�
 �	��	 �	��	 �	�� 		
   C  h8      ] "  )4L O 4LL O 5>  "  G  "���"���     w       g  filenamef  system/vm/inspect.scm�
 �	��	 �	��	 �	
��	 �	��	# �	
��	/ �	
��	/ �	�� 		3
   C       h   u   - 1 3 E   m       g  k
			 g  args			  g  filenamef  system/vm/inspect.scm�
 �	��	 �	�� 			
   C         h�  �  ]A45 HHHHHH4>  "  G  4>  "  G  44	5>  "  G  4
4	5>  "  G  44	5>  "  G  44	5>  "  G  KO  Q 4>  "  G  4>  "  G  44	5>  "  G  KO  Q 4>  "  G  4>  "  G  44	5>  "  G  KO  Q 4>  "  G  4>  "  G  44	5>  "  G  KO  Q 4>  "  G  4>  "  G  44	5>  "  G  KO Q 4>  "  G  4>  "  G  44	5>  "  G  44	5>  "  G  K O 6     �      g  x
	� g  commands	� g  quit		� g  print		� g  write		� g  display		� g  disassemble		� g  help		� g  c		 � g  c	 �! g  c	,� g  c	�� g  c	�> g  c	H�  g  filenamef  system/vm/inspect.scm�
	R
��		]	��		]	��		a	�� �	e	��,	i	���	m	���	q	��H	y	��� �	��� �	�� 	�  g  nameg  inspect� CRC    q       g  m
		,  g  filenamef  system/vm/inspect.scm�		
���	"
��:	*
���	4
���	R
�� 	�
   C6 