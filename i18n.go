//
// Copyright (c) 2019 Ted Unangst <tedu@tedunangst.com>
//
// Permission to use, copy, modify, and distribute this software for any
// purpose with or without fee is hereby granted, provided that the above
// copyright notice and this permission notice appear in all copies.
//
// THE SOFTWARE IS PROVIDED "AS IS" AND THE AUTHOR DISCLAIMS ALL WARRANTIES
// WITH REGARD TO THIS SOFTWARE INCLUDING ALL IMPLIED WARRANTIES OF
// MERCHANTABILITY AND FITNESS. IN NO EVENT SHALL THE AUTHOR BE LIABLE FOR
// ANY SPECIAL, DIRECT, INDIRECT, OR CONSEQUENTIAL DAMAGES OR ANY DAMAGES
// WHATSOEVER RESULTING FROM LOSS OF USE, DATA OR PROFITS, WHETHER IN AN
// ACTION OF CONTRACT, NEGLIGENCE OR OTHER TORTIOUS ACTION, ARISING OUT OF
// OR IN CONNECTION WITH THE USE OR PERFORMANCE OF THIS SOFTWARE.

package main

import (
    "net/http"
    "regexp"
    "log"
	"humungus.tedunangst.com/r/webs/login"
)

type i18n struct {
    Home string
    Atme string
    First string
    Combos string
    Chatter string
    Tags string
    Events string
    Longago string
    Saved string
    Honkers string
    Hcfs string
    Account string
    Morestuff string
    Myhonks string
    About string
    Front string
    Funzone string
    Xzone string
    Help string
    Search string
    Login string
    SetLang string
    Pwd string
    Refresh string
    ScrollDown string

    Username string
    Logout string
    Aboutme string
    Slayout string
    NoImg string
    MentionAll string
    AppleMapsLinks string
    Reaction string
    LanguageLabel string
    UpdateSettings string
    ChangePwd string
    OldPwd string
    NewPwd string
    ChangePwdBtn string

    NewChatter string
    Target string
    Noise string
    Chonk string
    Attach string
    Attachment string
    ExtAttachment string

    Onts string

    Original string
    Convoy string
    InReplyTo string
    Time string
    Location string
    Image string
    Unbonk string
    Bonk string
    Nope string
    HonkBack string
    EvenMore string
    Zonk string
    Deack string
    Ack string
    Unsave string
    Save string
    Untagged string
    UntagMe string
    Edit string
    Mute string
    Duration string
    HonkingTime string
    MakeNoise string
    MoreOptions string
    Description string
    GonBHonked string
    Cancel string
    Start string
    AddTime string
    Checkin string
    Preview string
    Actions string

    Newhonk string

    Version string
    Memory string
    Uptime string
    Cputime string

    AddNewHonker string
    Name string
    SkipSub string
    AddHonkerBtn string
    Expand string
    Unsub string
    Delete string
    Resub string
    Notes string
    Flavor string
    Optional string

    MoarHonks string

    Import string
    Fetch string
    Honks string
    HonksFrom string

    FunZoneGreet string

    Hfcs string
    NewFilter string
    FilterName string
    Matches string
    WhoWhere string
    IncludeAudience string
    TextMatches string
    IsAnnounce string
    AnnounceOf string
    Action string
    Reject string
    SkipMedia string
    Hide string
    Collapse string
    Rewrite string
    Replace string
    Expiration string
    BanHammerBtn string
    Date string
    Who string
    Pardon string

}

func getLangCookie(r *http.Request) string {
    langCookie, err := r.Cookie("lang")
    if err != nil {
          return "en"
    }
    return langCookie.Value
}



func setLangCookie (w http.ResponseWriter, r *http.Request) {
  var lang string
  lang = r.FormValue("lang")
  var IsLetter = regexp.MustCompile(`^([a-z]+)$`).MatchString

  if !IsLetter(lang) {
    lang = "wrong" // so !=2 is triggered
    if debugMode {
        log.Printf("lang cookie value is not letters")
    }
  }

  if len(lang) != 2 {
    if debugMode {
        log.Printf("lang cookie value is too long or too short. defaulting to eng")
    }
    lang = "en"
  }

  maxage := 3600 * 24 * 30 * 12
  if !debugMode {
       http.SetCookie(w, &http.Cookie{
            Name:     "lang",
            Value:    lang,
            MaxAge:   maxage,
            Secure:   true,
            HttpOnly: false,
       })
  } else {
       http.SetCookie(w, &http.Cookie{
            Name:     "lang",
            Value:    lang,
            MaxAge:   maxage,
            Secure:   false,
            HttpOnly: false,
       })
  }


  u := login.GetUserInfo(r)
  if u == nil {
    http.Redirect(w, r, "/", http.StatusSeeOther)
  }
}

func setLangStr (lang string) interface{} {
   switch lang {
       case "fr" :
           tlStr := i18n{
                "accueil",
                "mentions",
                "premier (first)",
                "combos",
                "discutaille",
                "balises",
                "événements",
                "il y a longtemps",
                "sauvegardés",
                "klaxonneurs",
                "filtrer (hcfs)",
                "compte",
                "plus de choses",
                "profil",
                "à propos",
                "tout le réseau connu",
                "zone fun",
                "récup",
                "aide",
                "rechercher",
                "connexion",
                "changer langue",
                "mot de passe",
                "rafraîchir",
                "bas de page",

                "nom d'utilisateur",
                "déconnexion",
                "à propos de moi",
                "plus mince",
                "pas d'images",
                "mentionner tout le monde",
                "liens apple maps",
                "réaction",
                "langage",
                "enregistrer paramètres",
                "changer le mot de passe",
                "ancien mdp",
                "nouveau mdp",
                "mettre à jour le mot de passe",

                "nouvelle discussion",
                "cible",
                "bruit",
                "et zlou",
                "joindre",
                "pièce jointe",
                "pièce jointe externe",

                "ontologies d'intérêts",

                "original",
                "convoi",
                "en réponse à",
                "temps",
                "lieu",
                "image",
                "départager",
                "partager",
                "nop",
                "klaxonner en retour",
                "encore plus",
                "suppr",
                "marquer comme non lu",
                "marquer comme lu",
                "oublier",
                "s'en souvenir",
                "débalisé",
                "me débaliser",
                "éditer",
                "sourdine",
                "durée",
                "il est temps de klaxonner",
                "faisons un peu d'bruit",
                "plus d'options",
                "description",
                "ça sera klaxonné",
                "annuler",
                "commence",
                "ajouter temps",
                "localiser",
                "aperçu",
                "Actions",

                "klaxonner",

                "version",
                "mémoire",
                "actif depuis",
                "cputime",

                "ajouter un klaxonneur",
                "pseudo",
                "ne pas suivre",
                "ajouter klaxonneur",
                "déplier",
                "se désabonner",
                "supprimer",
                "resuivre",
                "notes",
                "type",
                "optionnel",

                "un klaxon et peut être plus",

                "importer",
                "récupérer",
                "klaxons",
                "de",

                "Bienvenue dans la zone fun !",

                "Système de censure et filtrage de honk",
                "nouveau filtre",
                "nom du filtre",
                "correspondances",
                "qui ou où",
                "inclure l'audience",
                "texte contient",
                "est annoncé",
                "annonce à partir de",
                "mesures",
                "rejeter",
                "omettre médias",
                "cacher",
                "replier",
                "réécrire",
                "par",
                "expiration",
                "imposez votre volonté",
                "Date",
                "Qui",
                "pardonner",
           }
           return tlStr
       default:
          tlStr := i18n{
                "home",
                "@me",
                "first",
                "combos",
                "chatter",
                "tags",
                "events",
                "long ago",
                "saved",
                "honkers",
                "filters",
                "account",
                "more stuff",
                "my honks",
                "about",
                "front",
                "funzone",
                "xzone",
                "help",
                "search",
                "login",
                "set lang",
                "password",
                "refresh",
                "scroll down",

                "username",
                "logout",
                "about me",
                "skinny layout",
                "omit images",
                "mention all",
                "apple maps links",
                "reaction",
                "language",
                "update settings",
                "change password",
                "oldpass",
                "newpass",
                "update password",

                "new chatter",
                "target",
                "noise",
                "chonk",
                "attach",
                "Attachment",
                "External Attachment",

                "ontologies of interest",

                "original",
                "convoy",
                "in reply to ",
                "Time",
                "Location",
                "Image",
                "unbonk",
                "bonk",
                "nope",
                "honk back",
                "even more",
                "zonk",
                "deack",
                "ack",
                "unsave",
                "save",
                "untagged",
                "untag me",
                "edit",
                "mute",
                "duration",
                "it's honking time",
                "let's make some noise",
                "more options",
                "description",
                "it's gonna be honked",
                "cancel",
                "start",
                "add time",
                "checkin",
                "preview",
                "Actions",

                "new honk",

                "version",
                "memory",
                "uptime",
                "cputime",

                "add a new honker",
                "name",
                "skip sub",
                "add honker",
                "expand",
                "unsub",
                "delete",
                "resub",
                "notes",
                "flavor",
                "optional",

                "one honk maybe more",

                "import",
                "fetch",
                "honks",
                "from",

                "Welcome to the fun zone !",

                "Honk Filtering and Censorship System",
                "new filter",
                "filter name",
                "match",
                "who or where",
                "include audience",
                "text matches",
                "is announce",
                "announce of",
                "action",
                "reject",
                "skip media",
                "hide",
                "collapse",
                "rewrite",
                "replace",
                "expiration",
                "impose your will",
                "Date",
                "Who",
                "pardon",
          }
          return tlStr
   }
}
