package usrp

const (
	STRING_SOCKET_FAILURE     = "Socket failure"
	STRING_PRIVATE            = "Private"
	STRING_GROUP              = "Group"
	STRING_WINDOWS_PORT_REUSE = "On Windows, ignore the port reuse"

	// USRP Packet Types
	USRP_TYPE_VOICE       = 0
	USRP_TYPE_DTMF        = 1
	USRP_TYPE_TEXT        = 2
	USRP_TYPE_PING        = 3
	USRP_TYPE_TLV         = 4
	USRP_TYPE_VOICE_ADPCM = 5
	USRP_TYPE_VOICE_ULAW  = 6

	// TLV tags
	TLV_TAG_BEGIN_TX   = 0
	TLV_TAG_AMBE       = 1
	TLV_TAG_END_TX     = 2
	TLV_TAG_TG_TUNE    = 3
	TLV_TAG_PLAY_AMBE  = 4
	TLV_TAG_REMOTE_CMD = 5
	TLV_TAG_AMBE_49    = 6
	TLV_TAG_AMBE_72    = 7
	TLV_TAG_SET_INFO   = 8
	TLV_TAG_IMBE       = 9
	TLV_TAG_DSAMBE     = 10
	TLV_TAG_FILE_XFER  = 11
)
