# plan/ · Main (Zyklus 1)

Eine Datei pro Schritt. Der **Ausgang** eines Schritts ist wörtlich die **Zulassung** des nächsten.
Ein Schritt gilt erst als fertig, wenn alle Ausgang-Häkchen gesetzt sind, vorher wird der nächste nicht begonnen.

Alle Schritte sind ausgeplant. Was ein Schritt Neues lehrt, wird in den folgenden Dateien nachgezogen, bevor sie starten.

```
plan/
├── P01-fundament.md              8–12 Prompts   Fundament
├── P02-homepage.md              12–20 Prompts   Homepage
├── P03-accounts.md               8–12 Prompts   Accounts
├── P04-profil.md                  5–8 Prompts   Profil
├── P05-auth-tests.md              4–6 Prompts   Auth-Tests · MESSPUNKT 1
├── P06-anzeigen-basis.md         8–12 Prompts   Anzeigen · Basis
├── P07-anzeigen-verwalten.md     8–12 Prompts   Anzeigen · Verwalten
├── P08-anzeigen-tests.md          4–6 Prompts   Anzeigen-Tests · MESSPUNKT 2
├── P09-gruppen-basis.md          8–12 Prompts   Gruppen · Basis
├── P10-pinnwand.md               8–12 Prompts   Pinnwand
├── P11-code-review.md             3–6 Prompts   Code-Review
├── P12-haerten.md               12–20 Prompts   Härten & Mobile
├── P13-migrationen.md             4–6 Prompts   Migrationen
├── P14-deployment.md            10–15 Prompts   Deployment · MESSPUNKT 3
├── P15-linting-ci.md              3–5 Prompts   Linting & CI
└── P16-abnahme.md                 4–8 Prompts   Abnahme · MAIN FERTIG
```

**Summe Main:** 109–172 Prompts · bei 20–30 Prompts/Tag etwa 1,5–2 Wochen.

**Regeln**
- Ein P = ein Commit.
- Häkchen mit **(du)** prüfst und setzt du, alle anderen setzt Claude nur mit Nachweis (Build, Test, Antwort).
- Der nächste Schritt beginnt erst nach deiner Freigabe.
- Design-Wünsche gesammelt in einem Prompt.
- Neue Ideen landen in `draft.md`, nicht im laufenden Schritt.
