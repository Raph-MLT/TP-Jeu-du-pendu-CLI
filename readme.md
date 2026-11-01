Jeu du pendu en Go
Raphaël Maulat et Vivien Chaudron.
Le projet
C'est le jeu du pendu, en version terminal. Le programme choisit un mot au hasard et le joueur doit le deviner lettre par lettre, avant d'avoir épuisé ses essais.
Le jeu garde un historique de toutes les parties terminées. On peut aussi mettre une partie en pause et la reprendre plus tard.
Le projet est écrit en Go, uniquement avec la bibliothèque standard. Il sert à mettre en pratique la gestion des erreurs, la lecture et l'écriture de fichiers, et les formats CSV et JSON.
Lancer le jeu
Depuis le dossier du projet :

Plain Text
go run . jouer       # nouvelle partie
go run . reprendre   # reprendre la partie sauvegardée
go run . score       # voir l'historique des parties


Règles
Le mot est caché : chaque lettre est remplacée par un _. Une lettre est révélée au départ.
À chaque tour, on propose une lettre. Si elle est dans le mot, elle apparaît partout où elle se trouve. Sinon, on perd un essai.
On peut aussi tenter le mot entier. Un mauvais mot coûte 2 essais.
On a 6 essais au total.
Taper !save sauvegarde la partie et quitte le jeu.
Les fichiers du jeu
mot.txt : la liste des mots possibles, un par ligne.
save.json : la partie en pause. Il est créé par !save et supprimé quand on reprend la partie.
scores.csv : l'historique des parties, créé automatiquement à la fin de la première partie.