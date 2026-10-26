# P03 · Accounts

**Status:** ausgeplant · **Geschätzte Prompts:** 8–12

## Zulassung
_Ausgang von P02, unverändert übernommen_

- [ ] Die Startseite zeigt Hero, Manifest, Vorschau Anzeigen und Vorschau Gruppen.
- [ ] Impressum, Datenschutz und 404 sind erreichbar und im Footer verlinkt.
- [ ] Bei 375 px Breite gibt es kein horizontales Scrollen.
- [ ] (du) Design auf dem eigenen Handy abgenommen.
- [ ] Commit „P02 Homepage“.
- [ ] (du) Freigabe: P02 fertig.

## Ziel
Registrieren, Einloggen, Ausloggen – sicher und ohne Framework.

## Schritte (konkrete Prompts)
1. „schema.sql: Tabellen users (id, email unique, name unique, pw_hash, created_at) und sessions (token, user_id, expires_at).“
2. „auth.go: Registrieren und Login mit argon2id, Session-Token zufällig, HttpOnly-/Secure-/SameSite-Cookie, Logout löscht die Session in der DB.“
3. „CSRF-Token in allen Formularen; Formular ohne gültiges Token → 403.“
4. „Rate-Limit für Login: max. 5 Versuche pro Minute und IP. Fehlermeldung verrät nicht, ob die E-Mail existiert.“
5. „Templates login.html und register.html im Brand-Stil, Fehler am Feld, Labels immer sichtbar.“
6. Fix-Runde nach eigenem Durchklicken.

## Skills
go-error-handling · go-defensive (Login darf nichts verraten) · golang-popular-libraries (argon2-Paket prüfen).

## Hilfe (machst du selbst)
- [ ] (du) Einmal selbst registrieren, aus- und wieder einloggen, ein falsches Passwort probieren.

## Ausgang
_Wird 1:1 die Zulassung von P04._

- [ ] Registrieren, Login und Logout funktionieren im Browser.
- [ ] Passwörter liegen nur als argon2id-Hash in der DB.
- [ ] Logout entfernt die Session aus der DB.
- [ ] Formulare ohne CSRF-Token werden mit 403 abgelehnt.
- [ ] Ab dem 6. Login-Versuch pro Minute und IP kommt 429.
- [ ] (du) Selbst registriert, aus- und eingeloggt, falsches Passwort probiert.
- [ ] Commit „P03 Accounts“.
- [ ] (du) Freigabe: P03 fertig.
