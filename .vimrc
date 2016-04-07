filetype off                  " required

" set the runtime path to include Vundle and initialize
set rtp+=~/.vim/bundle/Vundle.vim
call vundle#begin()

Plugin 'xolox/vim-misc'
Plugin 'ekalinin/Dockerfile.vim'
"Plugin 'xolox/vim-easytags'
Plugin 'thinca/vim-quickrun'
Plugin 'Shougo/unite.vim'
Plugin 'Shougo/vimproc.vim'
Plugin 'Shougo/vimshell.vim'
Plugin 'nathanaelkane/vim-indent-guides'
Plugin 'regedarek/ZoomWin'
Plugin 'kien/rainbow_parentheses.vim'
Plugin 'corntrace/bufexplorer'
Plugin 'vim-ruby/vim-ruby'
Plugin 'vimoutliner/vimoutliner'
Plugin 'tpope/vim-speeddating'
Plugin 'itchyny/calendar.vim'
Plugin 'HerringtonDarkholme/vim-worksheet'
Plugin 'chrisbra/csv.vim'
Plugin 'vim-scripts/SearchComplete'
Plugin 'Lokaltog/vim-easymotion'
Plugin 'elzr/vim-json'
"Plugin 'lukaszkorecki/CoffeeTags'
Plugin 'mklabs/grunt.vim'
Plugin 'kchmck/vim-coffee-script'
Plugin 'digitaltoad/vim-jade.git'
Plugin 'vim-scripts/Conque-Shell'
Plugin 'kien/ctrlp.vim'
Plugin 'vim-scripts/dbext.vim'
Plugin 'Raimondi/delimitMate'
Plugin 'sickill/vim-monokai'
Plugin 'scrooloose/nerdcommenter'
Plugin 'scrooloose/nerdtree.git'
Plugin 'chrisbra/NrrwRgn'
Plugin 'aklt/plantuml-syntax'
Plugin 'ktvoelker/sbt-vim.git'
Plugin 'scrooloose/syntastic'
Plugin 'mkitt/tabline.vim.git'
Plugin 'godlygeek/tabular'
Plugin 'majutsushi/tagbar'
Plugin 'edkolev/tmuxline.vim'
"Plugin 'sirver/ultisnips'
Plugin 'vim-scripts/vcscommand.vim'
Plugin 'bling/vim-airline'
Plugin 'fatih/vim-go'
Plugin 'plasticboy/vim-markdown'
Plugin 'nelstrom/vim-markdown-folding'
Plugin 'terryma/vim-multiple-cursors'
Plugin 'derekwyatt/vim-scala'
Plugin 'mhinz/vim-signify'
Plugin 'jpalardy/vim-slime.git'
Plugin 'honza/vim-snippets'
Plugin 'tpope/vim-surround'
Plugin 'spiroid/vim-ultisnip-scala'
Plugin 'gmarik/Vundle.vim.git'
Plugin 'othree/xml.vim'
Plugin 'ervandew/supertab'
Plugin 'rking/ag.vim'
Plugin 'pangloss/vim-javascript'
Plugin 'tpope/vim-fugitive'
Plugin 'casecommons/vim-rails'
"Plugin 'Valloric/YouCompleteMe'

" All of your Plugins must be added before the following line
call vundle#end()            " required
filetype plugin indent on    " required

"execute pathogen#infect()
let g:ft_ignore_pat = '\.org'
filetype plugin indent on
filetype plugin on
au FileType php setl ofu=phpcomplete#CompletePHP
au FileType ruby,eruby setl ofu=rubycomplete#Complete
au FileType html,xhtml setl ofu=htmlcomplete#CompleteTags
au FileType c setl ofu=ccomplete#CompleteCpp
au FileType css setl ofu=csscomplete#CompleteCSS
autocmd BufNewFile,BufRead *.zsh setlocal filetype=sh

au! BufRead,BufWrite,BufWritePost,BufNewFile *.org
au BufEnter *.org            call org#SetOrgFileType()
command! OrgCapture :call org#CaptureBuffer()
command! OrgCaptureFile :call org#OpenCaptureFile()

syntax on

let g:utl_cfg_hdl_scm_http_system = "silent !google-chrome %u &"

set number
set smartindent
" size of tabstop
set tabstop=2
set incsearch

" size of an indent
set shiftwidth=2

" indent at beginning of a line
set smarttab

" always uses spaces instead of tab characters
set expandtab

nmap <F8> :TagbarToggle<CR>
let g:tagbar_autoshowtag = 0

nnoremap <F5> :silent update<Bar>silent !google-chrome %:p &<CR>
"nnoremap <F15>c :exe ':silent !google-chrome %'<CR>
"nmap <silent> <leader>w :exec 'silent !google-chrome % &'

let g:instant_markdown_slow = 1
let g:instant_markdown_autostart = 0

"let g:syntastic_disabled_filetypes=['scala', 'sbt']
let g:syntastic_scala_checkers=['']
"let g:syntastic_mode_map = { 'mode': 'passive', 'active_filetypes': [],'passive_filetypes': [] }
"nnoremap <C-w>E :SyntasticCheck<CR> :SyntasticToggleMode<CR>

let g:sqlutil_keyword_case = '\U'
let g:sqlutil_align_where = 1
let g:sqlutil_align_keyword_right = 0
let g:sqlutil_align_comma = 1
let g:multi_cursor_exit_from_insert_mode=0
let g:slime_target = "tmux"
let g:slime_paste_file = "$HOME/.slime_paste"
"set rtp+=$HOME/.local/lib/python2.7/site-packages/powerline/bindings/vim/

let g:airline#extensions#tabline#enabled = 1
let g:airline#extensions#tabline#fnamemod = ':t'
let g:airline_powerline_fonts = 1
" Always show statusline
set laststatus=2
set noshowmode

fun! <SID>StripTrailingWhitespaces()
  let l = line(".")
  let c = col(".")
  %s/\s\+$//e
  call cursor(l, c)
endfun
autocmd BufWritePre * :call <SID>StripTrailingWhitespaces()

"" make YCM compatible with UltiSnips (using supertab)
"let g:ycm_key_list_select_completion = ['<C-n>', '<Down>']
"let g:ycm_key_list_previous_completion = ['<C-p>', '<Up>']
"let g:SuperTabDefaultCompletionType = '<C-n>'

"" better key bindings for UltiSnipsExpandTrigger
"let g:UltiSnipsExpandTrigger = "<tab>"
"let g:UltiSnipsJumpForwardTrigger = "<tab>"
"let g:UltiSnipsJumpBackwardTrigger = "<s-tab>"

let g:dbext_default_profile_cm_main_20141110 = 'type=PGSQL:user=postgres:passwd=postgres:host=172.16.9.54:port=6432:dbname=cm_main_20141110'
let g:dbext_default_profile_cm_holy= 'type=PGSQL:user=postgres:passwd=postgres:host=172.16.9.54:port=6432:dbname=cm_holy'
let g:dbext_default_profile_cm_org_test= 'type=PGSQL:user=postgres:passwd=postgres:host=172.16.9.54:port=6432:dbname=cm_org_test'
let g:dbext_default_profile_cm_main_test= 'type=PGSQL:user=postgres:passwd=postgres:host=172.16.9.54:port=6432:dbname=load_data_test'


let g:ctrlp_working_path_mode = 'a'

syntax enable
set t_Co=256
colorscheme monokai
"let g:solarized_termcolors=256
"colorscheme solarized
"if has('gui_running')
"  set background=light
"else
"  "set background=dark
"  set background=light
"endif

let g:signify_vcs_list = [ 'svn', 'git' ]
let g:ycm_server_keep_logfiles = '1'
set t_ut=

if exists(":Tabularize")
  nmap <Leader>a= :Tabularize /=<CR>
  vmap <Leader>a= :Tabularize /=<CR>
  nmap <Leader>a: :Tabularize /:\zs<CR>
  vmap <Leader>a: :Tabularize /:\zs<CR>
endif

inoremap <silent> <Bar>   <Bar><Esc>:call <SID>align()<CR>a

function! s:align()
  let p = '^\s*|\s.*\s|\s*$'
  if exists(':Tabularize') && getline('.') =~# '^\s*|' && (getline(line('.')-1) =~# p || getline(line('.')+1) =~# p)
    let column = strlen(substitute(getline('.')[0:col('.')],'[^|]','','g'))
    let position = strlen(matchstr(getline('.')[0:col('.')],'.*|\s*\zs.*'))
    Tabularize/|/l1
    normal! 0
    call search(repeat('[^|]*|',column).'\s\{-\}'.repeat('.',position),'ce',line('.'))
  endif
endfunction

if filereadable(glob(".vimrc.local"))
  source .vimrc.local
endif

let g:calendar_google_calendar = 1
let g:calendar_google_task = 1

let maplocalleader = ',,'

let g:EclimCompletionMethod = 'omnifunc'
let g:SuperTabDefaultCompletionType = 'context'

" Rainbow parantheses
let g:rbpt_colorpairs = [
    \ ['brown',       'RoyalBlue3'],
    \ ['Darkblue',    'SeaGreen3'],
    \ ['darkgray',    'DarkOrchid3'],
    \ ['darkgreen',   'firebrick3'],
    \ ['darkcyan',    'RoyalBlue3'],
    \ ['darkred',     'SeaGreen3'],
    \ ['darkmagenta', 'DarkOrchid3'],
    \ ['brown',       'firebrick3'],
    \ ['gray',        'RoyalBlue3'],
    \ ['black',       'SeaGreen3'],
    \ ['darkmagenta', 'DarkOrchid3'],
    \ ['Darkblue',    'firebrick3'],
    \ ['darkgreen',   'RoyalBlue3'],
    \ ['darkcyan',    'SeaGreen3'],
    \ ['darkred',     'DarkOrchid3'],
    \ ['red',         'firebrick3'],
    \ ]
let g:rbpt_max = 16
let g:rbpt_loadcmd_toggle = 0
au VimEnter * RainbowParenthesesToggle
au Syntax * RainbowParenthesesLoadRound
au Syntax * RainbowParenthesesLoadSquare
au Syntax * RainbowParenthesesLoadBraces

" Automatically reload vim config(s)
augroup myvimrc
  au!
  au BufWritePost .vimrc,_vimrc,vimrc,.gvimrc,_gvimrc,gvimrc so $MYVIMRC | if has('gui_running') | so $MYGVIMRC | endif
augroup END

let indent_guides_auto_colors = 0
autocmd VimEnter,Colorscheme * :hi IndentGuidesOdd  guibg=darkgrey   ctermbg=236
autocmd VimEnter,Colorscheme * :hi IndentGuidesEven guibg=darkgrey   ctermbg=237
let indent_guides_color_change_percent = 10
let indent_guides_guide_size = 2

" VimShell
" Use current directory as vimshell prompt.
let g:vimshell_prompt_expr =
\ 'escape(fnamemodify(getcwd(), ":~").">", "\\[]()?! ")." "'
let g:vimshell_prompt_pattern = '^\%(\f\|\\.\)\+> '

" easytags
let g:easytags_async = 1

