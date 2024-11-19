# Notes


### O que aprendi?

general - commands
- comando cat le conteúdo de arquivos, posso contenar resultados de arquivos ou até mesmo juntar eles em um só.
- echo "conteúod" > arquivo sobreescreve o conteúdo do arquivo com a string passada >> apenas faz um append.
- ! o terminal entende como um sinal de expansão dentro de uma string, você pode evitar esse sinal usando single quotes ''

git 
- git possuí escopo de sistema, global, local, working tree cada escopo sobreescreve o git config de cada.
- git tree meio de guardar diretórios do git
- git blob meio de guardar arquivos do git
- git branch é um ponteiro para o último commit hash da branch
- caso de uso base do git rebease, escreve no git history da base branch commits na ponta, 
rebase não gera um novo commit. Rebase reescreve o histórico de commits.
- git log -P é o equivalente para o --no-pager(adicionei um alias para o git -P ser default)
- git log --oneline mostra um resumo do histórico
- git log --graph constroi o grafo de commits da quela branch
- git log --parents incluí o commit hash que originou aquele commit(útil para branchs com muitos merges)
- git config --list, lista todas as configs.
- git guardas as informações no config por meio de chave e valor, git permite existir chave dúplicadas com valore diferentes
- git config --get pega valor de uma chave
- git guara o conteúdo dos arquivos em binário, para ler o conteúdo é necessário ler por meio de um parse utf-8
- git cat-file é um command que faz o parse do conteúdo de um binário para você, foneça a hash de um blob.
- git reset --soft hash desfaz o commit e mantém as alterações staged(git add .)
- git reset --hard hash desfaz o commit e deleta permanentemente as alterações
- git remote add {nome_remote} {path} é o comando que define um repositório como remote(fonte da verdade), geralmente chamado de origin.
path pode ser um paht local para um repositório ou um link do github de um repositório hospedado por um produto git.
- .gitignore possuío vários patterns, *{algo} ignora todos algo, !{} negação do ignore, 
por exemplo *.txt \n !vamogremio.txt ele irá ignorar todos arquivos txt exceto vamogremio.txt.
/{alog} irá ignorar apenas no mesmo nível de diretório do .gitignore. Para mais pattrerns busque .gitignore patterns documentation
- git reflog é o comando que guarda todas as ações feitas em uma branch, alterações não apenas de commits e sim de checkouts, pulls, reset e delete.

## Git reset with a commit in the middle of the branch
I have a branch in this branch i have this cenário of commits when a use the command
-> git log --oneline
output: 
hash-3 C 
hash-2 B 
hash-1 A

I would like to remove the commit B, so I would do git reset --hard hash-2, but when I do this, git has removed the commit with message C, why this haapend?

GitHub
The command git reset --hard hash-2 resets your branch to the state of commit hash-2, which means the commits on top of hash-2 (in this case, commit C) are discarded. If you want to remove commit B without losing commit C, you should use an interactive rebase instead:

Run git rebase -i HEAD~3.
In the editor that appears, delete the line corresponding to commit B.
Save and close the editor.
This will remove commit B and keep commits A and C.

## Git recover lost changes
git reflog é o comando que guarda todas as ações feitas em uma branch
```
git reflog (find the commit sha at HEAD@{1})
git cat-file -p <commit sha>
git cat-file -p <tree sha>
git cat-file -p <blob sha> > slander.md
git add .
git commit -m "B: recovery"
```

Or we can just do `git merge <commitish>`
o commitish pode ser `(branch, tag, commit, HEAD@{1})`
Então se nós executarmo `git reflog` pegar o ` HEAD@{1}` da mudanças que desejamos e executar
`git merge  HEAD@{1}` nós iremos recuperar as mudanças do commit + o commit merge.