# Notes

general - commands
- comando cat le conteúdo de arquivos, posso contenar resultados de arquivos ou até mesmo juntar eles em um só.
- echo "conteúod" > arquivo sobreescreve o conteúdo do arquivo com a string passada >> apenas faz um append.
- ! o terminal entende como um sinal de expansão dentro de uma string, você pode evitar esse sinal usando single quotes ''

#### git 
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
- git reset --soft hash, a hash que você passa é o index aonde o git irá parar de desfazer as mudanças, desfaz o commit e mantém as alterações staged(git add .)
Exemplo:
hash-A - 1
hash-B - 2
hash-C - 3
`git reset --soft hash-C` irá desfazer os commits A e B, `git reset --soft hash-B` irá desfazer o commit A, `git reset --soft hash-A` não irá fazer nada pois, hash-A ja é sua HEAD.
- git reset --hard hash desfaz o commit e deleta permanentemente as alterações
- git remote add {nome_remote} {path} é o comando que define um repositório como remote(fonte da verdade), geralmente chamado de origin.
path pode ser um paht local para um repositório ou um link do github de um repositório hospedado por um produto git.
- .gitignore possuío vários patterns, *{algo} ignora todos algo, !{} negação do ignore, 
por exemplo *.txt \n !vamogremio.txt ele irá ignorar todos arquivos txt exceto vamogremio.txt.
/{alog} irá ignorar apenas no mesmo nível de diretório do .gitignore. Para mais pattrerns busque .gitignore patterns documentation
- git reflog é o comando que guarda todas as ações feitas em uma branch, alterações não apenas de commits e sim de checkouts, pulls, reset e delete.
- Use `git restore --staged .` to remove things from staged area(`git add .`), `git restore .` remove any change not staged
- git revert will remove commits by creating a new commit to this new change, keeping the history of changes with out removing commits only moving forward. Use `git revert commit-hash` to revert until this specific commit
- Use `git diff` to see the differences between commits and state of my code
Examples:
```bash
# show the changes between the working tree and the last commit
git diff

# show the change between two commits
git diff COMMIT_HASH_1 COMMIT_HASH_2

# show the differences between the previous commit and the current state, including the last commit and uncommitted changes
git diff HEAD~1
```
- Use `git cherry-pick {desire-commit-hash}` to get a specific commit from a branch and insert in you working branch

- Use `git tag -a "<version>" -m "<description version>"` to create a tag
- Use `git push origin --tags` to publish in the remote repository the local tags
- Use `git tag` to list all local tags


## Git clean
The command git clean -f -q -- /Users/olaisaac/projects/repos/src/megacorp/partner list is used to remove untracked files from the working directory. Here is a breakdown of the command:

git clean: This is the Git command used to remove untracked files from the working directory.
-f: This flag stands for "force". It is required to actually remove the files. Without this flag, git clean will not do anything.
-q: This flag stands for "quiet". It suppresses the output of the command, so you won't see a list of the files that are being removed.
--: This is used to separate the command options from the file path. It ensures that the file path is interpreted correctly.
/Users/olaisaac/projects/repos/src/megacorp/partner list: This is the specific file or directory that you want to clean. In this case, it is the partner list file located in the specified path.
So, this command will forcefully and quietly remove the untracked file partner list from the specified directory.

## How to stash

- Do some changes, no need to stage then
- `git stash`git will save this changes on a index area where he will sync with your HEAD(last commit)
- Use `git stash list` to see all stash that u have
- Use `git stash pop` to get the most recent stash in the index stack(LIFO) and apply those changes into your working directory.
- Use `git stash drop` will remove most recent stash
- Use `git stash clear` will clean up every element of the stack.


## How to Squash

Perhaps confusingly, squashing is done with the git rebase command! Here are the steps to squash the last n commits:

Start an interactive rebase with the command git rebase -i HEAD~n, where n is the number of commits you want to squash.
Git will open your default editor with a list of commits. Change the word pick to squash for all but the first commit.
Save and close the editor.
The -i flag stands for "interactive," and it allows us to edit the commit history before Git applies the changes. HEAD~n is how we reference the last n commits. HEAD points to the current commit (as long as we're in a clean state) and ~n means "n commits before HEAD."

In the list of commits the commits that you change to squash will be those commits that will be rebased onto the first commit with pick.

## Git reset with a commit in the middle of the branch
I have a branch in this branch i have this cenário of commits when a use the command
-> git log --oneline
output: 
hash-3 C 
hash-2 B 
hash-1 A

I would like to remove the commit B, so I would do git reset --hard hash-2, but when I do this, git has removed the commit with message C, why this haapend?

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

## Git bisect

That's where the git bisect command comes in. Instead of manually checking all the commits (O(n) for Big O nerds), 
git bisect allows us to do a binary search (O(log n) for Big O nerds) to find the commit that introduced the bug.

How to bisect 
1. Start the bisect with `git bisect start`
2. Select a "good" commit with `git bisect good <commitish>` (a commit where you're sure the bug wasn't present)
3. Select a bad commit via `git bisect bad <commitish>` (a commit where you're sure the bug was present)
4. Git will checkout a commit between the good and bad commits for you to test to see if the bug is present
5. Execute `git bisect good` or `git bisect bad` to say the current commit is good or bad
6. Loop back to step 4 (until git bisect completes)
7. Exit the bisect mode with `git bisect reset`

During the state of `git bisect` your current working directory will be at the commit to choose bad or good, 
so you can run you tests or run your local system to tests de commit

You can execute a script where exit 1 or 0 will say if the analised commit is bad(1) ou good(0).
Run `git bisect start`
Mark a commit as bad
Mark a commit as good
Automate the rest of the bisect with `git bisect run <some-scricpt.sh>`


## What Is a Worktree?
A worktree (or "working tree" or "working directory") is just the directory on your filesystem where the code you're tracking with Git lives. 
Usually, it's just the root of your Git repo (where the .git directory is). It contains:
- Tracked files (files that Git knows about)
- Untracked files (files that Git doesn't know about)
- Modified files (files that Git knows about that have been changed since the last commit)