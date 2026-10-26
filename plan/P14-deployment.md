# P14 · Deployment · MESSPUNKT 3

**Status:** ausgeplant · **Geschätzte Prompts:** 10–15

## Zulassung
_Ausgang von P13, unverändert übernommen_

- [ ] `schema.sql` ist ersetzt durch `migrations/`.
- [ ] `migrate.go` führt nur neue Migrationen aus.
- [ ] Vorhandene Testdaten sind nach dem Umstieg vollständig da.
- [ ] Commit „P13 Migrationen“.
- [ ] (du) Freigabe: P13 fertig.

## Ziel
Live unter HTTPS, mit Backup, das nachweislich funktioniert.

## Schritte (konkrete Prompts)
1. „deploy/Caddyfile für {domain} mit automatischem HTTPS, Proxy auf :8080.“
2. „deploy/lab.service: systemd, Restart=always, Umgebung aus /etc/lab.env.“
3. „deploy/backup.sh + Timer: täglicher pg_dump, 14 Tage aufbewahren.“
4. „Fehler-Mail: bei Panic oder mehr als 10 Antworten mit 5xx in 5 Minuten, über SMTP aus der Umgebung.“
5. „deploy/README.md: Schritt für Schritt vom leeren Server bis online.“
6. „Restore-Test: Backup in leere DB einspielen, Zeilenzahlen vergleichen.“

## Skills
golang-continuous-integration (Vorbereitung) · go-logging

## Hilfe (machst du selbst)
- [ ] (du) Server oder VPS bereitstellen, Domain und DNS einrichten.
- [ ] (du) SMTP-Zugang für die Fehler-Mail besorgen.

## Ausgang
_Wird 1:1 die Zulassung von P15._

- [ ] Die Seite läuft unter der Domain mit HTTPS.
- [ ] systemd startet den Dienst nach einem Absturz neu.
- [ ] Das tägliche Backup läuft, ein Restore in eine leere DB ist erfolgreich getestet.
- [ ] (du) Eine Test-Fehler-Mail ist angekommen.
- [ ] Commit „P14 Deployment“.
- [ ] (du) Freigabe: P14 fertig.
