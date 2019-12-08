package molekcheatez

//go:generate stringer -type=PuzzleID

//PuzzleID is the id number of a puzzle
type PuzzleID int32

//Puzzle IDs
const (
	InvalidPuzzleID PuzzleID = iota
	DEPRECATED_Tutorial1
	HydrogenPeroxide
	DEPRECATED_UnbondingTutorial
	NitrousOxide
	Ethanol
	GHB
	Cyanamide
	Tranylcypromine
	MDMA
	Dimercaprol
	Nicotine
	DEPRECATED_Fentanyl
	Ketamine
	Propofol
	DMT
	Caffeine
	Mebicar
	Phencyclidine
	Amphetamine
	Barbital
	Methylhexanamine
	Metformin
	Sertraline
	Aspirin
	Modafinil
	Epinephrine
	Benzylpiperazine
	Pyrazinamide
	Lidocaine
	DEPRECATED_RCS4
	KojicAcid
	DMSO
	DimethylEther
	Piracetam
	Chloroform
	Valnoctamide
	Propranolol
	Diazepam
	Dapsone
	Mescaline
	Dimethylthiambutene
	SolitaireMain
	SwitchToBonus
	SwitchToMain
	Placeholder
	Ibuprofen
	LevelEditor
	Propylthiouracil
	Isoflurane
	Inositol
	PyridoxalPhosphate
	SulfurHexafluoride
	Hydrochlorothiazide
	Mianserin
	Fluorouracil
	Cytophosphane
	Ethambutol
	Nitroglycerin
	Fosfomycin
	SulfurMustard
	Sarin
	Paracetamol
	ClodronicAcid
	Lamotrigine
	Methaqualone
	Hexestrol
	Thiotepa
	Mexiletine
	Promazine
	Altretamine
	Pyrimethamine
	Lysine
	VitaminC
	Fluoperazine
	Neramexane
	ChloralHydrate
	Ethchlorvynol
	Thalidomide
	Bithionol
	SolitaireBonus //oh?
)

//go:generate stringer -type=PrecursorID

//PrecursorID is the id number of a precursor
type PrecursorID int32

//Precursor IDs
const (
	InvalidPrecursor PrecursorID = iota
	Water
	Ammonia
	HydrochloricAcid
	SulfuricAcid
	Acetone
	Cyclohexane
	Methanol
	EthyleneGlycol
	Propene
	AceticAcid
	Benzene
	Carbamide
	PhosphoricAcid
	HydrofluoricAcid
	Hydrazine
	Butanone
	ThionylChloride
	Dioxane
	Toluene
	FormicAcid
	Butylene
	Isobutane
	Triazine
)

//go:generate stringer -type=Instruction

//Instruction is one action the emitters can do
type Instruction uint8

//Instructions
const (
	NoInstruction Instruction = iota
	SlideLeft
	SlideRight
	Push
	Pull
	RotateLeft
	RotateRight
	AddHydrogen
	RemoveHydrogen
	Delete
	Output
	ShuntHydrogen
)


