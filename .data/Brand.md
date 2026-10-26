# Brand.md — LAB (Arbeitstitel)

> **Status:** Aus den vorhandenen Assets abgeleitet. Das Logo ist ein JPEG ohne Transparenz und ohne lesbare Wortmarke. Die Bildinhalte wurden programmatisch analysiert (Palette, Histogramm, Geometrie, OCR); eine visuelle Bildprüfung war in dieser Session nicht möglich. Semantische Deutungen des Emblems sind deshalb bewusst offen und in „Offene Punkte“ markiert.

---

## 1. Design Read

| Feld | Wert |
|---|---|
| Artefakt | Landing Page / erste Homepage |
| Zielgruppe | Kreative, Kuratierende, mögliche Partner (zu bestätigen) |
| Bildsprache | Editorial Organic — dunkles Oliv, Papiercreme, Bronze, ein sparsamer Limetten-Akzent |
| Modus | Greenfield |
| Visual Variance | 7 / 10 — asymmetrisch, großzügig, nicht Template-förmig |
| Motion Intensity | 5 / 10 — leise Reveals, präzise Hover-Zustände |
| Information Density | 4 / 10 — luftig, ein Fokus pro Viewport |
| Asset Dependence | 8 / 10 — Emblem und Bildmaterial tragen die Marke |
| Brand Fidelity | 9 / 10 — Olive/Creme/Gold/Limette und Emblem bleiben unangetastet |

**Leitidee:** Eine ruhige, kuratierte Galerie mit Labor-Charakter. Kein lauter SaaS-Look, keine Tech-Farben. Die Seite soll wirken wie ein sauber gesetztes Print-Magazin, das sich im Browser bewegt.

---

## 2. Asset-Inventar

| Datei | Format / Größe | Hintergrund | Verwendung |
|---|---|---|---|
| `logo-dark.jpeg` | JPEG, 1600 × 1600 | Schwarz `#000000` | Emblem auf dunklen Flächen |
| `logo-dark1.jpeg` | JPEG, 1600 × 1600 | Schwarz `#000000` | Alternative dunkle Fassung, etwas tiefer im Bild |
| `logo-white.jpeg` | JPEG, 1600 × 1600 | Creme `#ece5d8` | Emblem auf hellen / Papier-Flächen |
| `default-anzeige-svg.svg` | SVG, 800 × 450 | Oliv-Verlauf | 16:9-Platzhalter für Galerie / Anzeigen |
| `default-avatar.svg` | SVG, 128 × 128 | Oliv-Verlauf | Runder Profil-Platzhalter |

### Emblem

- Zentrales, rund/oval gefasstes Emblem mit symmetrischer, organisch-ornamentaler Binnenstruktur.
- Keine lesbare Wortmarke im Bild; OCR liefert keinen verwertbaren Text.
- Inhaltsfläche `logo-dark.jpeg`: ca. `x 260–1340`, `y 372–1324` bei 1600 × 1600. Das Emblem nimmt rund 67 % der Breite ein.
- Das Emblem ist kein freigestelltes Asset. Für Web-Nutzung gibt es zwei saubere Wege:
  1. **Sofort:** `mix-blend-mode: screen` auf dunklen Flächen, `multiply` auf Creme-Flächen. Die JPEG-Hintergründe verschwinden dadurch optisch.
  2. **Besser:** transparente SVG-/PNG-Fassung vom Original erstellen lassen.
- Schutzraum: mindestens 25 % der Emblemhöhe an allen Seiten.
- Mindestgröße: 40 px Höhe für das Emblem, 24 px nur als rein dekorativer Stempel.

### Platzhalter

- `default-anzeige-svg.svg` ist inhaltlich ein **Beispielbild**: „BEISPIELANZEIGE“ / „Galerie – Bildmaterial folgt“ mit Mikrofon-Linienicon.
- `default-avatar.svg` ist eine **Silhouette**, kein echtes Porträt.
- Beide dürfen als ehrliche Platzhalter eingesetzt werden, aber nicht als echte Kundenreferenzen oder echte Personen ausgegeben werden.

---

## 3. Marke

| Feld | Stand |
|---|---|
| Name | **LAB** (Arbeitstitel aus dem Projektordner, zu bestätigen) |
| Claim | offen |
| Kategorie | Kreativlabor / Galerie / Plattform (zu bestätigen) |
| Versprechen | Kuratiert, handwerklich, ruhig, mit Substanz |
| Persönlichkeit | Erdend, präzise, zurückhaltend, geheimnisvoll, hochwertig |
| Tonalität | Kurze deutsche Sätze. Konkret statt werblich. Keine Superlative, keine erfundenen Zahlen. |
| Anti-Positionierung | Kein Neon-Tech, kein Purple-Gradient, keine Emoji-Deko, keine generischen SaaS-Karten |

**Markenformel:** Dunkles Oliv als Bühne, Papiercreme als Denkfläche, Bronze als Linie, Limette als Funke.

---

## 4. Farben

### Freigegebene Palette (2026-10-26) · gilt vor allem Folgenden

Gemessen aus `hero1.jpeg` (Leder + Gold-Schriftzug), `hero2.jpg` (Neon-Kassette) und dem Gold-Schriftzug. Von Roman freigegeben („das bin ich“). Die ältere Oliv/Sage-Palette weiter unten ist damit **abgelöst**. Tokens in `tailwind.config.js`.

| Gruppe | Token | Hex | Einsatz |
|---|---|---|---|
| Grund | `ink` | `#0b0c0e` | Hintergrund: kühles Anthrazit (seit 2026-10-26 statt Leder-Schwarz `#110b04`, das zu braun wirkte) |
| Leder | `leather-dark` | `#2a1f15` | Flächen |
| Leder | `leather` | `#4d4135` | Karten, Linien |
| Leder | `paper` | `#efe6cf` | Text (Creme) |
| Gold | `gold-shine` | `#f5e6b2` | Highlights, Glanz im Button-Verlauf |
| Gold | `gold-light` | `#e1cf9a` | Akzenttext |
| Gold | `gold` | `#cab682` | Buttons, Linien |
| Gold | `gold-brass` | `#9a8553` | Messing, Schatten |
| Grün | `spray` | `#6a8652` | dezenter Grün-Akzent |
| Grün | `cassette` | `#213a24` | grüne Flächen |
| Grün | `neon` | `#52e158` | nur als Funke (Label, Leuchtpunkt), nie als Fläche |
| Grün | `neon-light` | `#8dff6e` | heller Neon-Touch mit Glow: Krone, aktiver Menüpunkt, Labels, Naht unter dem Hero, Hover |

Buttons: Gold-Verlauf `#f5e6b2 → #cab682 → #9a8553`, Text `#1a1206`.

### Rohpalette aus den Assets (abgelöst)

| Token | Hex | Herkunft / Rolle |
|---|---|---|
| `ink` | `#090802` | Dominantes Dunkel aus den Logos |
| `black` | `#000000` | Logo-Hintergrund, maximaler Kontrast |
| `olive-950` | `#18170d` | Tiefe Flächen, Verläufe |
| `olive-900` | `#1e200b` | Panels, Footer |
| `olive-800` | `#2a2816` | Karten, Eingabefelder |
| `olive-700` | `#3d3b1f` | Erhabene Flächen, Verlauf |
| `olive-600` | `#484423` | Rahmen, Badge |
| `olive-500` | `#5e5420` | Sekundärflächen |
| `olive-400` | `#686236` | Hairlines auf Dunkel |
| `sage-300` | `#9e9c61` | Gedämpfter Text, Meta |
| `sage-200` | `#b2b28c` | Sekundärtext |
| `sage-100` | `#d6dcab` | Helle Sage-Flächen |
| `cream-50` | `#f4eee2` | Hellster Papier-Ton |
| `cream-100` | `#ece5d8` | Papier / Primärtext auf Dunkel |
| `cream-200` | `#e6e1be` | Weicher Text |
| `gold-300` | `#d7af7e` | Warmes Highlight |
| `gold-500` | `#a8905e` | Bronze / Linien / Labels |
| `gold-600` | `#8b6a2f` | Tiefe Bronze |
| `lime-300` | `#cdf7b3` | Weicher Akzent |
| `lime-400` | `#acf99b` | Hauptakzent |
| `lime-500` | `#8add73` | Hover / aktiver Zustand |
| `green-600` | `#67a94e` | Sekundärer Grün-Akzent |
| `green-700` | `#559132` | Tiefes Grün |

### Semantische Tokens

```css
:root {
  --bg: #090802;
  --bg-elevated: #18170d;
  --surface: #1e200b;
  --surface-2: #2a2816;
  --paper: #ece5d8;
  --paper-bright: #f4eee2;
  --ink: #18170d;
  --text: #ece5d8;
  --text-muted: #b2b28c;
  --text-subtle: #9e9c61;
  --line: rgba(236, 229, 216, 0.14);
  --line-strong: rgba(236, 229, 216, 0.28);
  --accent: #cdf7b3;
  --accent-strong: #acf99b;
  --accent-deep: #8add73;
  --gold: #a8905e;
  --gold-deep: #8b6a2f;
  --focus: #cdf7b3;
}
```

### Verteilung

- **60 %** dunkles Oliv/Schwarz (`bg`, `surface`)
- **25 %** Papiercreme (`paper`) für ruhige Inhaltsflächen
- **10 %** Bronze/Sage für Linien, Labels, Sekundärtext
- **5 %** Limette für Interaktion und einen einzigen visuellen Funken

### Kontrast (geprüft)

| Kombination | Verhältnis | Einsatz |
|---|---:|---|
| `#ece5d8` auf `#090802` | 16,0 : 1 | Primärtext |
| `#b2b28c` auf `#090802` | 9,2 : 1 | Sekundärtext |
| `#cdf7b3` auf `#090802` | 16,8 : 1 | Akzenttext |
| `#a8905e` auf `#090802` | 6,5 : 1 | Bronze-Text / Linien |
| `#8b6a2f` auf `#090802` | 4,0 : 1 | nur große Schrift / dekorativ |
| `#18170d` auf `#ece5d8` | 15,1 : 1 | Text auf Papier |
| `#3d3b1f` auf `#ece5d8` | 9,1 : 1 | Text auf Papier, weich |
| `#686236` auf `#090802` | 3,2 : 1 | nur Hairlines, nie Fließtext |

**Regeln:** Keine neuen Blau-, Violett- oder Pinktöne. Limette niemals als große Fläche. Bronze ist Linie und Stimme, nicht Hintergrund.

---

## 5. Typografie

### Paarung

| Rolle | Schrift | Gewichte | Begründung |
|---|---|---|---|
| Display / Headlines | **Instrument Serif** | 400, 400 italic | Editorial, galeriehaft, mit eigenem Charakter; kein Standard-Serif-Look |
| Text / UI | **Karla** | 400, 500, 600, 700 | Warme Grotesk mit leicht organischer Note, sehr gute Lesbarkeit |
| Labels / Meta | **Karla** uppercase | 600 | Kein drittes Schriftsystem nötig |

**Nicht verwenden:** Inter, Roboto, Arial, System-UI als Display, Fraunces.

### Skala

| Token | Größe | Line-height | Tracking | Einsatz |
|---|---|---|---|---|
| `display-xl` | `clamp(3.75rem, 9.5vw, 8.5rem)` | 0.90 | -0.03em | Hero, ein Satz |
| `display-l` | `clamp(3rem, 6.5vw, 6rem)` | 0.94 | -0.025em | große Section-Headlines |
| `h2` | `clamp(2.25rem, 4vw, 3.5rem)` | 1.00 | -0.02em | Abschnittstitel |
| `h3` | `clamp(1.5rem, 2.2vw, 2rem)` | 1.10 | -0.01em | Karten / Untertitel |
| `body-l` | `18px` (mobile `16px`) | 1.65 | 0 | Intro, Lead |
| `body` | `16px` | 1.65 | 0 | Fließtext |
| `small` | `13px` | 1.50 | 0.01em | Meta, Captions |
| `kicker` | `11px` | 1.40 | 0.18em | Uppercase-Labels |
| `button` | `14px` | 1.00 | 0.02em | CTAs |

### Verteilung

- Headlines **linksbündig**, nie zentriert als Default.
- Hero-Headline: maximal 2–3 Zeilen, 7 von 12 Spalten.
- Lead-Text: maximal 58–68 Zeichen pro Zeile.
- Section-Intros: 5 von 12 Spalten, rechts viel Luft.
- Kicker immer mit kurzer Hairline davor: `— 01 / GALERIE`.
- Serif nur für Headlines und große Zitate. Karla für alles Funktionale.
- Keine fette Serif. Hierarchie entsteht über Größe und Weißraum, nicht über Gewicht.

---

## 6. Layout & Raster

| Token | Wert |
|---|---|
| Spalten | 12 |
| Max Width | `1280px` |
| Gutter | `clamp(16px, 2vw, 32px)` |
| Außenabstand | `clamp(20px, 5vw, 72px)` |
| Section Padding | `clamp(80px, 10vw, 160px)` |
| Baseline | 8 px |
| Radien | `2px` (tight), `8px` (inputs), `20px` (media/cards), `999px` (pills/avatars) |

### Kompositionsregeln

- Ein klarer Fokus pro Viewport.
- Asymmetrie statt 50/50: bevorzugt **7/5**, **8/4** oder **5/7**.
- Große Flächen dürfen leer sein. Leere ist Teil der Marke.
- Bilder und Emblem bekommen Luft, keine dekorativen Rahmen.
- Trennungen über Hairlines und Abstand, nicht über Schatten.
- Mobile: einspaltig, aber Typo-Kontrast bleibt großzügig.

### Breakpoints

`640 / 768 / 1024 / 1280 / 1536`

---

## 7. UI-System

### Buttons

| Typ | Style | Zustände |
|---|---|---|
| Primary | Papiercreme gefüllt, `ink` Text, Pill-Form, Pfeil im Kreis | Hover: Limette; Active: 1 px nach unten; Focus: 2 px Limette-Ring |
| Secondary | transparent, 1 px Bronze-Hairline, Creme-Text | Hover: Bronze-Fläche; Focus: Limette-Ring |
| Text-Link | Karla 600, 1 px Unterlinie mit Offset | Hover: Limette, Unterlinie zieht sich |
| Icon-Button | 44 × 44 px, runder Rahmen | Hover: Fläche `surface-2` |

- Mindesthöhe 44 px.
- Kein Drop-Shadow. Tiefe entsteht über Fläche und Hairline.
- Pfeil-Icon 16 px, `stroke-width: 1.5`.

### Karten / Galerie

- Radius 20 px, 1 px `--line`, kein Schatten.
- Hover: Bild skaliert `1.02`, Hairline wird Bronze, Label wandert 4 px.
- 16:9 für Galerie, 4:5 für Porträts, 1:1 für Emblem-Kacheln.
- Caption immer außerhalb des Bildes, klein, uppercase, Sage.

### Navigation

- Höhe 72 px, transparent über dem Hero, ab Scrollen `backdrop-filter: blur(16px)` + `bg/80`.
- Links: Logo, 3–4 Hauptlinks, CTA rechts.
- Mobile: vollflächiges Menü in `olive-950`, Links in Serif, große Touch-Flächen.
- Aktiver Link: Limette-Punkt, keine fette Schrift.

### Formulare

- Eingabefeld: `8px` Radius, `surface-2`, 1 px Hairline.
- Focus: 2 px Limette-Ring, kein Glow.
- Fehler: Bronze-Text + Hairline, kein Rot (nicht in der Palette).
- Labels immer sichtbar, nie nur Placeholder.

### Tags / Badges

- Pill, 1 px Hairline, `kicker`-Typo.
- Erlaubt: „01 — Galerie“, „Neu“, „Limitiert“.
- Keine Emoji-Badges.

### Iconografie

- Linienicons, 24 px, `stroke-width: 1.5`, runde Enden.
- Vorhandenes Mikrofon-Icon als Referenz für den Stil.
- Keine Icon-Fonts, keine Emoji als Ersatz.
- Fehlendes Icon: ehrlicher Platzhalter `▢` oder Label, nicht selbst gezeichnet.

### Textur & Atmosphäre

- Feines Filmkorn als Overlay, 3–4 % Opazität.
- Radialer Olive-Glow hinter dem Hero, sehr subtil.
- Keine Purple-Gradients, keine Neon-Schatten.
- `text-wrap: pretty` für Fließtext, `text-wrap: balance` für Headlines.

---

## 8. Motion

| Token | Wert |
|---|---|
| Easing | `cubic-bezier(0.22, 1, 0.36, 1)` |
| Reveal | 600 ms, `opacity 0 → 1`, `translateY(16px) → 0` |
| Stagger | 70 ms |
| Hover | 200 ms |
| Button Arrow | 300 ms |
| Page Load | Hero-Elemente 0 / 80 / 160 ms versetzt |

- Kein Bounce, kein Spring, kein Dauer-Pulsieren.
- `prefers-reduced-motion`: alle Transforms aus, nur Opacity.
- Parallax maximal 40 px, nur auf Desktop.

---

## 9. Logo-Nutzung

| Fläche | Datei | Technik |
|---|---|---|
| Dunkel (`#090802`, `#18170d`) | `logo-dark.jpeg` | `mix-blend-mode: screen` |
| Creme (`#ece5d8`, `#f4eee2`) | `logo-white.jpeg` | `mix-blend-mode: multiply` |
| Unebene Fotografie | keins | Logo nur auf ruhiger Scrim-Fläche |

- Emblem nie verzerren, nie einfärben, nie anschneiden.
- Schutzraum: 25 % der Emblemhöhe.
- Mindestgröße: 40 px.
- Als Favicon: quadratischer Ausschnitt des Emblems, 64 × 64, auf `#090802`.
- Sobald eine transparente SVG-Fassung vorliegt: Blend-Modes entfernen und SVG einsetzen.
- **Wortmarke fehlt:** Name „LAB“ wird typografisch in Instrument Serif gesetzt, bis eine echte Wortmarke vorliegt.

---

## 10. Tailwind-Token

Für Tailwind v4 in `styles.css`:

```css
@import "tailwindcss";

@theme {
  --color-ink: #090802;
  --color-surface: #18170d;
  --color-surface-2: #1e200b;
  --color-surface-3: #2a2816;
  --color-paper: #ece5d8;
  --color-paper-bright: #f4eee2;
  --color-cream: #e6e1be;
  --color-sage: #b2b28c;
  --color-sage-deep: #9e9c61;
  --color-olive: #3d3b1f;
  --color-olive-deep: #484423;
  --color-gold: #a8905e;
  --color-gold-deep: #8b6a2f;
  --color-lime: #cdf7b3;
  --color-lime-strong: #acf99b;
  --color-lime-deep: #8add73;

  --font-display: "Instrument Serif", Georgia, serif;
  --font-sans: "Karla", "Helvetica Neue", sans-serif;

  --radius-tight: 2px;
  --radius-input: 8px;
  --radius-card: 20px;
  --radius-pill: 999px;

  --spacing-section: clamp(5rem, 10vw, 10rem);
  --container-content: 80rem;
}
```

**Font-Loading:**

```html
<link rel="preconnect" href="https://fonts.googleapis.com">
<link rel="preconnect" href="https://fonts.gstatic.com" crossorigin>
<link href="https://fonts.googleapis.com/css2?family=Instrument+Serif:ital@0;1&family=Karla:wght@400;500;600;700&display=swap" rel="stylesheet">
```

---

## 11. Inhaltsleitlinien

- Deutsch, ruhig, konkret. Keine austauschbaren Agenturfloskeln.
- Headline = ein Gedanke, kein Buzzword-Stapel.
- Keine erfundenen Statistiken, Kundenlogos oder Testimonials.
- Fehlende Inhalte werden als Platzhalter markiert: `[OFFER]`, `[AUDIENCE]`, `[PROOF]`, `[CTA]`.
- Jede Sektion muss eine Aufgabe haben: Orientierung, Beweis, Vertrauen, Handlung.
- Bilder nur mit echtem Material oder dem vorhandenen `default-anzeige`-Platzhalter.

### Copy-Struktur der Landing Page

1. **Hero:** Emblem, Name, ein Satz Positionierung, ein CTA.
2. **Manifest:** 2–3 kurze Absätze, warum es LAB gibt.
3. **Galerie / Auswahl:** 3–6 Karten mit `default-anzeige` als Platzhalter.
4. **Leistung / Angebot:** 3 klare Punkte, keine Feature-Wand.
5. **Über / Haltung:** ein ruhiger Textblock, 5 Spalten.
6. **CTA:** eine einzige Handlung, keine konkurrierenden Buttons.
7. **Footer:** Kontakt, Impressum, Datenschutz, Social.

---

## 12. Offene Punkte

1. **Markenname und Claim:** Ist „LAB“ korrekt? Gibt es eine Wortmarke?
2. **Angebot:** Galerie, Marktplatz, Studio, Musikplattform oder etwas anderes?
3. **Zielgruppe:** Kreative, Käufer:innen, Partner, B2B?
4. **Primäre Handlung:** Warteliste, Kontakt, Demo, Galerie ansehen?
5. **Logo-Semantik:** Das Emblem ist ornamental-kreisförmig; bitte bestätigen, was es darstellt (z. B. Sonne, Lyra, Baum, Wappen).
6. **Transparentes Logo:** Liegt eine SVG-/PNG-Fassung ohne Hintergrund vor? Falls nein, arbeiten wir zunächst mit Blend-Modes.
7. **Echte Bilder:** Gibt es Galerie-, Produkt- oder Teamfotos? Sonst bleiben die SVG-Platzhalter sichtbar.
8. **Postgres:** Braucht die erste Landing Page bereits eine Datenbank (z. B. Warteliste) oder reicht die statische Auslieferung?
