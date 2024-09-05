package entry

const otpSuffix = "otp"

type OtpEntry struct {
	fileEntry
}

func NewOtpEntry(parent *DirEntry, fileName string) *OtpEntry {
	return &OtpEntry{
		fileEntry{
			parent:   parent,
			fileName: fileName,
		},
	}
}
