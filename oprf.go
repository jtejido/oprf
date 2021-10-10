package oprf

const (
	version         = "VOPRF07-"
	seedDST         = "Seed-"
	challengeDST    = "Challenge-"
	finalizeDST     = "Finalize-"
	compositeDST    = "Composite-"
	hashToGroupDST  = "HashToGroup-"
	hashToScalarDST = "HashToScalar-"
)

// Mode specifies properties of the OPRF protocol.
type Mode = uint8

// https://www.ietf.org/archive/id/draft-irtf-cfrg-voprf-07.html#section-3
const (
	Base       Mode = 0x00
	Verifiable Mode = 0x01
)

type CipherSuiteID = uint16

const (
	// OPRFP256 represents the OPRF with Ristretto-255 and SHA-512.
	OPRFRistretto SuiteID = 0x0001
	// OPRFP256 represents the OPRF with Decaf-448 and SHAKE-256.
	OPRFDecaf SuiteID = 0x0002
	// OPRFP256 represents the OPRF with P-256 and SHA-256.
	OPRFP256 SuiteID = 0x0003
	// OPRFP384 represents the OPRF with P-384 and SHA-512.
	OPRFP384 SuiteID = 0x0004
	// OPRFP521 represents the OPRF with P-521 and SHA-512.
	OPRFP521 SuiteID = 0x0005
)
