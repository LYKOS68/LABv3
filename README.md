Bereich	Technologie
Frontend =	Next.js 15 (SSR/CSR), Tailwind CSS, shadcn/ui, Framer Motion, Mapbox GL JS
API	= Go (REST, JSON, Middleware, Auth, Rate Limiting)
Datenbank =	Postgres + PostGIS (alle Daten + Geo), Redis (Sessions, Rate Limits)
Hintergrundprozesse	 = Go Worker (Inserat ablaufen lassen, E-Mail senden, Boost-Reset)
Events	= Postgres LISTEN/NOTIFY oder Redis Pub/Sub → Go Consumer
Tracking	= PostHog (self-hosted, Events + Analytics)
Werbung	= Eigene Boost-Logik in Go + Postgres (kein externes Ad-Netzwerk)
Sicherheit	= JWT Auth, bcrypt, HTTPS, CSRF-Token, Input-Validierung in Go
Datenschutz	= DSGVO-Export in Go, Soft Delete, Resend für transaktionale Mails

Unsere Endplattform ist ein Online-Kleinanzeigenmarkt für die Musikbranche, auf dem Nutzer Anzeigen erstellen, durchsuchen und verwalten können. Die Kernfunktionen sind Inserate für Beats, Mix & Master, Musikproduktionen, Ghostwriting, Musikvideos und Coverdesign, Kategorien, Suche und Filter nach Stadt und Umkreis, Preise, Nutzerprofile sowie die direkte Kontaktaufnahme zwischen Interessenten und Anbietern. Ergänzend bietet die Plattform Gruppen, Community-Beiträge, Kurse und den Austausch zwischen Musikern, Produzenten und Dienstleistern. und zuerst eine vollständige Funktionslandkarte brauchen: Frontend, Navigation, Inserate, Werbung, Tracking, Event-Queues, Bidding, Slot-Kommunikation, Standort, Nutzerkonto, Nachrichten, APIs, Sicherheit, Datenschutz und Fehlerbehandlung. Jeder Bereich muss in konkrete Unterfunktionen zerlegt und gegen die untersuchten Quellen abgeglichen werden – ohne unbelegte Implementierungsdetails als Fakten auszugeben. 

