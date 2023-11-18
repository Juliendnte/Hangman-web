# Hangman-web

Bienvenue dans un site reprenant le jeu du pendu. Les règles sont simples, il faut juste deviner une lettre qui pourrait être dans le mot mis avec des underscores. Pour jouer à ce jeu, tu devras tout d'abord mettre un pseudo, puis choisir un des 12 modes proposés (multilettres est le rassemblement de mots en 3 lettres jusqu'à mot en + de 10 lettres et impossible est une liste de 300 000 mots du dictionnaire).
Ensuite, il va falloir deviner une des lettres du mot prises dans le mode choisi si tu réussis à deviner le mot, tu tomberas sur une page résultat qui te montrera ton score et une proposition de relancer. Sinon, si tu n'arrives pas à deviner le mot et que tu vois le dessin d'un pendu alors tu ne pourras pas rejouer contemplant ton score.
Mais nous avons d'autres fonctionnalités, lors de ta dernière vie, un indice sur une lettre prise au hasard du mot a deviné sera proposé.


Nous avons choisi le thème du taro pour l'aspect graphique du site lorsque tu arrives à la 12e carte, tu auras perdu, car derrière cette carte se cache un pendu.
Au niveau du code notre arboréance est simple nous avons 6 dossiers :

- assets, conserve le css et les images du site.
- controller, pour sa part appelle les fonctions du jeu et a tout le mechanisme pour que le site marche convenablement, ses fonctions seront appellés lorsque l'utilisateur sera sur une des pages du site. Il y a aussi, les routes treatment qui elles ne seront pas affichées.
- Hangman, comme son nom l'indice contient toutes les fonctions du Hangman.
- mot est le dossier ayant toutes la liste de mots possible.
- routeur, gère les routes et le maintient du serveur.
- temp, possède les html des pages et un fichier go pour leurs initialisations.