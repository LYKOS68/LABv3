# P13 · Migrationen

**Status:** ausgeplant · **Geschätzte Prompts:** 4–6

## Zulassung
_Ausgang von P12, unverändert übernommen_

- [ ] Alle Formulare validieren serverseitig und zeigen den Fehler am Feld.
- [ ] Honeypot und Post-Limit sind aktiv.
- [ ] Keine Seite scrollt bei 375 px horizontal.
- [ ] (du) Mobile-Design abgenommen.
- [ ] Commit „P12 Härten“.
- [ ] (du) Freigabe: P12 fertig.

## Ziel
Datenbank-Änderungen sicher machen, bevor echte Daten kommen.

## Schritte (konkrete Prompts)
1. „Vorher: pg_dump der lokalen DB nach backups/vor-p13.sql.“
2. „migrations/0001_init.sql aus dem aktuellen schema.sql, Tabelle schema_migrations für den Stand.“
3. „migrate.go: führt beim Start nur noch nicht gelaufene Migrationen in einer Transaktion aus. schema.sql löschen.“
4. „Zeigen, dass alle Testdaten noch da sind und go test grün ist.“

## Skills
go-error-handling

## Hilfe (machst du selbst)
- [ ] (du) Vorher einen Dump der lokalen DB ziehen.

## Ausgang
_Wird 1:1 die Zulassung von P14._

- [ ] `schema.sql` ist ersetzt durch `migrations/`.
- [ ] `migrate.go` führt nur neue Migrationen aus.
- [ ] Vorhandene Testdaten sind nach dem Umstieg vollständig da.
- [ ] Commit „P13 Migrationen“.
- [ ] (du) Freigabe: P13 fertig.
