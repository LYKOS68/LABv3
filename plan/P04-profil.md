# P04 · Profil

**Status:** ausgeplant · **Geschätzte Prompts:** 5–8

## Zulassung
_Ausgang von P03, unverändert übernommen_

- [ ] Registrieren, Login und Logout funktionieren im Browser.
- [ ] Passwörter liegen nur als argon2id-Hash in der DB.
- [ ] Logout entfernt die Session aus der DB.
- [ ] Formulare ohne CSRF-Token werden mit 403 abgelehnt.
- [ ] Ab dem 6. Login-Versuch pro Minute und IP kommt 429.
- [ ] (du) Selbst registriert, aus- und eingeloggt, falsches Passwort probiert.
- [ ] Commit „P03 Accounts“.
- [ ] (du) Freigabe: P03 fertig.

## Ziel
Öffentliches Profil, eigenes Profil bearbeiten, geschützte Bereiche.

## Schritte (konkrete Prompts)
1. „schema.sql: profiles (user_id, rolle rapper/producer/engineer, bio max. 500 Zeichen, avatar_url nullable).“
2. „profile.go: GET /u/{name} öffentlich mit Rolle, Bio und Avatar-Platzhalter (default-avatar.svg).“
3. „GET/POST /me/profil: eigenes Profil bearbeiten, Validierung am Feld.“
4. „requireAuth-Middleware für alle /me-Routen, ohne Login Redirect auf /login?next=…, danach zurück.“
5. „Nav zeigt eingeloggt: Profil + Logout, ausgeloggt: Login + Registrieren.“
6. Fix-Runde nach dem Test mit zwei Accounts.

## Skills
go-defensive

## Hilfe (machst du selbst)
- [ ] (du) Zwei Test-Accounts anlegen und gegenseitig die Profile prüfen.

## Ausgang
_Wird 1:1 die Zulassung von P05._

- [ ] `/u/{name}` ist öffentlich sichtbar.
- [ ] Bearbeiten ist nur für das eigene Profil möglich.
- [ ] Alle `/me`-Routen leiten ohne Login auf `/login` um.
- [ ] (du) Mit zwei Accounts die Profile gegenseitig geprüft.
- [ ] Commit „P04 Profil“.
- [ ] (du) Freigabe: P04 fertig.
