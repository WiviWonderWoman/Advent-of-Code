Sammanfattning av algoritmen:

1. Parsa input - Läs in alla junction boxes med deras 3D-koordinater
2. Beräkna alla avstånd - För varje par av boxes, räkna ut avståndet mellan dem
3. Sortera efter avstånd - Kortaste avståndet först
4. Union-Find - En smart datastruktur som håller koll på vilka boxes som är ihopkopplade (samma krets). När du kopplar ihop två boxes slås deras kretsar ihop.
5. Gör 1000 kopplingar - Ta de 1000 kortaste kopplingarna och slå ihop kretsarna
6. Räkna kretsarnas storlekar - Se hur många boxes varje krets innehåller
7. Multiplicera de tre största - Svaret!

/ claude

