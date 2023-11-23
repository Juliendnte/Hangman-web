# Hangman-web

Bienvenue dans un site reprenant le jeu du pendu. Les règles sont simples, il faut juste deviner une lettre qui pourrait être dans le mot mis avec des underscores. 
Pour jouer à ce jeu, tu devras tout d'abord faire un git clone "https://github.com/Juliendnte/Hangman-web.git", puis faire un go run . ,et cliquer sur le lien affiché. Dans le jeu, il faudra ensuite mettre un pseudo, puis choisir un des 12 modes proposés (multilettres est le rassemblement de mots en 3 lettres jusqu'à mot en + de 10 lettres et dantesque une liste de 300 000 mots du dictionnaire).
Ensuite, il va falloir deviner une des lettres du mot pris dans le mode choisi si tu réussis à deviner le mot, tu tomberas sur une page résultat qui te montrera ton score et une proposition de relancer. Sinon, si tu n'arrives pas à deviner le mot et que tu vois le dessin d'un pendu. Alors tu contempleras ton score et ne pourras pas rejouer. Nous ne sommes pas méchants, on t'a laissé tout de même 12 pdv avant de mourir.
Mais, nous avons aussi d'autres fonctionnalités, lors de ta dernière vie, un indice sur une lettre prise au hasard du mot a deviné sera proposé une fois.


Nous avons choisi le thème du taro/médiéval pour l'aspect graphique du site lorsque tu arrives à la 12e carte, tu auras perdu, car derrière cette carte, se cache un pendu.
Au niveau du code notre arboréance est simple nous avons 6 dossiers :

- assets, qui conserve le css , les fonts ,et les images du site.
- controller, pour sa part contient le code source qui permet de déclarer les gestionnaires qui vont être associés aux multiples routes du serveur HTTP.
- Hangman, appelle toutes les fonctions du Hangman.
- mot est le dossier ayant toutes la liste de mots possible.
- routeur, gère le code source permettant d'initialiser le serveur HTTP, mais également l'initialisation de son arborescence de route avec leur gestionnaire associé.
- temp, possède les html des pages et un fichier go pour leurs initialisations.