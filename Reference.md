# Reference for Values

In the JSON files, most of the values are stored as integers. This file has a handy reference so you can know what to change values to.

All of these were found by decompiling the Molek-Syntez assembly except for the instructions.

## Puzzle IDs
```
0    InvalidPuzzleID
1	DEPRECATED_Tutorial1
2	HydrogenPeroxide
3	DEPRECATED_UnbondingTutorial
4	NitrousOxide
5	Ethanol
6	GHB
7	Cyanamide
8	Tranylcypromine
9	MDMA
10	Dimercaprol
11 	Nicotine
12 	DEPRECATED_Fentanyl
13 	Ketamine
14 	Propofol
15	DMT
16	Caffeine
17	Mebicar
18	Phencyclidine
19	Amphetamine
20	Barbital
12	Methylhexanamine
22	Metformin
23	Sertraline
24	Aspirin
25	Modafinil
26	Epinephrine
27	Benzylpiperazine
28	Pyrazinamide
29	Lidocaine
30	DEPRECATED_RCS4
31	KojicAcid
32	DMSO
33	DimethylEther
34	Piracetam
35	Chloroform
36	Valnoctamide
37	Propranolol
38	Diazepam
39	Dapsone
40	Mescaline
41	Dimethylthiambutene
42	SolitaireMain
43	SwitchToBonus
44	SwitchToMain
45	Placeholder
46	Ibuprofen
47	LevelEditor
48	Propylthiouracil
49	Isoflurane
50	Inositol
51	PyridoxalPhosphate
52	SulfurHexafluoride
53	Hydrochlorothiazide
54	Mianserin
55	Fluorouracil
56	Cytophosphane
57	Ethambutol
58	Nitroglycerin
59	Fosfomycin
60	SulfurMustard
61	Sarin
62	Paracetamol
63	ClodronicAcid
64	Lamotrigine
65	Methaqualone
66	Hexestrol
67	Thiotepa
68	Mexiletine
69	Promazine
70	Altretamine
71	Pyrimethamine
72	Lysine
73	VitaminC
74	Fluoperazine
75	Neramexane
76	ChloralHydrate
77	Ethchlorvynol
78	Thalidomide
79	Bithionol
80	SolitaireBonus  
```

## Precursors

```
0   InvalidPrecursor
1	Water
2	Ammonia
3	HydrochloricAcid
4	SulfuricAcid
5	Acetone
6	Cyclohexane
7	Methanol
8	EthyleneGlycol
9	Propene
10	AceticAcid
11	Benzene
12	Carbamide
13	PhosphoricAcid
14	HydrofluoricAcid
15	Hydrazine
16	Butanone
17	ThionylChloride
18	Dioxane
19	Toluene
20	FormicAcid
21	Butylene
22	Isobutane
23	Triazine
```

## Instructions
Instructions are stored in-game as a uint. Molek-Cheatez uses the uints internally, but displays them to you (in `String()` and JSON encoding) as characters for ease of use. Both are displayed here.

```
0  .    NoInstruction
1  >	SlideLeft
2  <	SlideRight
3  ^	Push
4  v	Pull
5  \	RotateLeft
6  /	RotateRight
7  +	AddHydrogen
8  -	RemoveHydrogen
9  x	Delete
10 #	Output
11 ~	ShuntHydrogen
```