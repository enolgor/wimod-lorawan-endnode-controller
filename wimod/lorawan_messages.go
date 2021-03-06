package wimod

import (
	"encoding/binary"
	"fmt"
)

// LORAWAN_MSG_ACTIVATE_DEVICE_REQ

type ActivateDeviceReq struct {
	wimodMessageImpl
	Address    uint32
	AppSessKey Key
	NwkSessKey Key
}

func NewActivateDeviceReq(address uint32, appSessKey Key, nwkSessKey Key) *ActivateDeviceReq {
	req := &ActivateDeviceReq{}
	req.Init()
	req.Address = address
	req.AppSessKey = appSessKey
	req.NwkSessKey = nwkSessKey
	return req
}

func (p *ActivateDeviceReq) Init() {
	p.code = LORAWAN_MSG_ACTIVATE_DEVICE_REQ
}

func (p *ActivateDeviceReq) String() string {
	return fmt.Sprintf("ActivateDeviceReq[Address: %08X, AppSessKey: %v, NwkSessKey: %v]", p.Address, p.AppSessKey, p.NwkSessKey)
}

func (p *ActivateDeviceReq) Encode() ([]byte, error) {
	buff := make([]byte, 4)
	binary.LittleEndian.PutUint32(buff, p.Address)
	buff = append(buff, EncodeKey(&p.NwkSessKey)...)
	buff = append(buff, EncodeKey(&p.AppSessKey)...)
	return buff, nil
}

// LORAWAN_MSG_ACTIVATE_DEVICE_RSP

type ActivateDeviceResp struct {
	wimodMessageStatusImpl
}

func NewActivateDeviceResp() *ActivateDeviceResp {
	resp := &ActivateDeviceResp{}
	resp.Init()
	return resp
}

func (p *ActivateDeviceResp) Init() {
	p.code = LORAWAN_MSG_ACTIVATE_DEVICE_RSP
}

func (p *ActivateDeviceResp) String() string {
	return fmt.Sprintf("ActivateDeviceResp[]")
}

func (p *ActivateDeviceResp) Decode(payload []byte) error {
	p.Status = payload[0]
	return lorawanStatusCheck(p.Status)
}

// LORAWAN_MSG_SET_JOIN_PARAM_REQ

type SetJoinParamReq struct {
	wimodMessageImpl
	AppEUI EUI
	AppKey Key
}

func NewSetJoinParamReq(appEUI EUI, appKey Key) *SetJoinParamReq {
	req := &SetJoinParamReq{}
	req.Init()
	req.AppEUI = appEUI
	req.AppKey = appKey
	return req
}

func (p *SetJoinParamReq) Init() {
	p.code = LORAWAN_MSG_SET_JOIN_PARAM_REQ
}

func (p *SetJoinParamReq) String() string {
	return fmt.Sprintf("SetJoinParamReq[AppEUI: %v, AppKey: %v]", p.AppEUI, p.AppKey)
}

func (p *SetJoinParamReq) Encode() ([]byte, error) {
	buff := EncodeEUI(&p.AppEUI)
	buff = append(buff, EncodeKey(&p.AppKey)...)
	return buff, nil
}

// LORAWAN_MSG_SET_JOIN_PARAM_RSP

type SetJoinParamResp struct {
	wimodMessageStatusImpl
}

func NewSetJoinParamResp() *SetJoinParamResp {
	resp := &SetJoinParamResp{}
	resp.Init()
	return resp
}

func (p *SetJoinParamResp) Init() {
	p.code = LORAWAN_MSG_SET_JOIN_PARAM_RSP
}

func (p *SetJoinParamResp) String() string {
	return fmt.Sprintf("SetJoinParamResp[]")
}

func (p *SetJoinParamResp) Decode(payload []byte) error {
	p.Status = payload[0]
	return lorawanStatusCheck(p.Status)
}

// LORAWAN_MSG_JOIN_NETWORK_REQ

type JoinNetworkReq struct {
	wimodMessageImpl
}

func NewJoinNetworkReq() *JoinNetworkReq {
	req := &JoinNetworkReq{}
	req.Init()
	return req
}

func (p *JoinNetworkReq) Init() {
	p.code = LORAWAN_MSG_JOIN_NETWORK_REQ
}

func (p *JoinNetworkReq) String() string {
	return fmt.Sprintf("JoinNetworkReq[]")
}

func (p *JoinNetworkReq) Encode() ([]byte, error) {
	return []byte{}, nil
}

// LORAWAN_MSG_JOIN_NETWORK_RSP

type JoinNetworkResp struct {
	wimodMessageStatusImpl
}

func NewJoinNetworkResp() *JoinNetworkResp {
	resp := &JoinNetworkResp{}
	resp.Init()
	return resp
}

func (p *JoinNetworkResp) Init() {
	p.code = LORAWAN_MSG_JOIN_NETWORK_RSP
}

func (p *JoinNetworkResp) String() string {
	return fmt.Sprintf("JoinNetworkResp[]")
}

func (p *JoinNetworkResp) Decode(payload []byte) error {
	p.Status = payload[0]
	return lorawanStatusCheck(p.Status)
}

// LORAWAN_MSG_JOIN_NETWORK_TX_IND

type JoinNetworkTxInd struct {
	wimodMessageStatusImpl
	ChannelIdx       byte
	DataRateIdx      byte
	NumTxPackets     byte
	TRXPowerLevel    byte
	RFMessageAirtime uint32
}

func NewJoinNetworkTxInd() *JoinNetworkTxInd {
	ind := &JoinNetworkTxInd{}
	ind.Init()
	return ind
}

func (p *JoinNetworkTxInd) Init() {
	p.code = LORAWAN_MSG_JOIN_NETWORK_TX_IND
}

func (p *JoinNetworkTxInd) String() string {
	return fmt.Sprintf("JoinNetworkTxInd[Status: 0x%02X, ChannelIdx: %d, DataRateIdx: %d, NumTxPackets: %d, TRXPowerLevel: %d, RFMessageAirtime: %d]", p.Status, p.ChannelIdx, p.DataRateIdx, p.NumTxPackets, p.TRXPowerLevel, p.RFMessageAirtime)
}

func (p *JoinNetworkTxInd) Decode(bytes []byte) error {
	p.Status = bytes[0]
	if p.Status != LORAWAN_MSG_JOIN_NETWORK_TX_IND_STATUS_OK && p.Status != LORAWAN_MSG_JOIN_NETWORK_TX_IND_STATUS_OK_ATTACHMENT {
		p.Status = LORAWAN_MSG_JOIN_NETWORK_TX_IND_STATUS_ERROR
	}
	if p.Status == LORAWAN_MSG_JOIN_NETWORK_TX_IND_STATUS_OK_ATTACHMENT {
		p.ChannelIdx = bytes[1]
		p.DataRateIdx = bytes[2]
		p.NumTxPackets = bytes[3]
		p.TRXPowerLevel = bytes[4]
		p.RFMessageAirtime = binary.LittleEndian.Uint32(bytes[5:9])
	}
	return nil
}

// LORAWAN_MSG_JOIN_NETWORK_IND

type JoinNetworkInd struct {
	wimodMessageStatusImpl
	Address     uint32
	ChannelIdx  byte
	DataRateIdx byte
	RSSI        byte
	SNR         byte
	RxSlot      byte
}

func NewJoinNetworkInd() *JoinNetworkInd {
	ind := &JoinNetworkInd{}
	ind.Init()
	return ind
}

func (p *JoinNetworkInd) Init() {
	p.code = LORAWAN_MSG_JOIN_NETWORK_IND
}

func (p *JoinNetworkInd) String() string {
	return fmt.Sprintf("JoinNetworkInd[Status: 0x%02X, Address: 0x%08X, ChannelIdx: %d, DataRateIdx: %d, RSSI: %d, SNR: %d, RxSlot: %d]", p.Status, p.Address, p.ChannelIdx, p.DataRateIdx, p.RSSI, p.SNR, p.RxSlot)
}

func (p *JoinNetworkInd) Decode(bytes []byte) error {
	p.Status = bytes[0]
	if p.Status != LORAWAN_MSG_JOIN_NETWORK_IND_STATUS_OK && p.Status != LORAWAN_MSG_JOIN_NETWORK_IND_STATUS_OK_ATTACHMENT {
		p.Status = LORAWAN_MSG_JOIN_NETWORK_IND_STATUS_ERROR
	}
	if p.Status == LORAWAN_MSG_JOIN_NETWORK_IND_STATUS_OK_ATTACHMENT {
		p.Address = binary.LittleEndian.Uint32(bytes[1:5])
		p.ChannelIdx = bytes[5]
		p.DataRateIdx = bytes[6]
		p.RSSI = bytes[7]
		p.SNR = bytes[8]
		p.RxSlot = bytes[9]
	}
	return nil
}

// LORAWAN_MSG_SEND_UDATA_REQ

type SendUDataReq struct {
	wimodMessageImpl
	Port    byte
	Payload []byte
}

func NewSendUDataReq(port byte, payload []byte) *SendUDataReq {
	req := &SendUDataReq{}
	req.Init()
	req.Port = port
	req.Payload = payload
	return req
}

func (p *SendUDataReq) Init() {
	p.code = LORAWAN_MSG_SEND_UDATA_REQ
}

func (p *SendUDataReq) String() string {
	return fmt.Sprintf("SendUDataReq[Port: %d, Payload: 0x%X]", p.Port, p.Payload)
}

func (p *SendUDataReq) Encode() ([]byte, error) {
	buff := []byte{p.Port}
	buff = append(buff, p.Payload...)
	return buff, nil
}

// LORAWAN_MSG_SEND_UDATA_RSP

type SendUDataResp struct {
	wimodMessageStatusImpl
	RemainingTime uint32
}

func NewSendUDataResp() *SendUDataResp {
	resp := &SendUDataResp{}
	resp.Init()
	return resp
}

func (p *SendUDataResp) Init() {
	p.code = LORAWAN_MSG_SEND_UDATA_RSP
}

func (p *SendUDataResp) String() string {
	return fmt.Sprintf("SendUDataResp[RemainingTime: %d]", p.RemainingTime)
}

func (p *SendUDataResp) Decode(payload []byte) error {
	p.Status = payload[0]
	switch p.Status {
	case LORAWAN_STATUS_OK:
		return nil
	case LORAWAN_STATUS_CHANNEL_BLOCKED:
		p.RemainingTime = binary.LittleEndian.Uint32(payload[1:5])
		return fmt.Errorf("LORAWAN_STATUS_CHANNEL_BLOCKED: Remaining Time: %d", p.RemainingTime)
	default:
		return lorawanStatusCheck(payload[0])
	}
}

// LORAWAN_MSG_SEND_UDATA_TX_IND

type SendUDataTxInd struct {
	wimodMessageStatusImpl
	ChannelIdx       byte
	DataRateIdx      byte
	NumTxPackets     byte
	TRXPowerLevel    byte
	RFMessageAirtime uint32
}

func NewSendUDataTxInd() *SendUDataTxInd {
	ind := &SendUDataTxInd{}
	ind.Init()
	return ind
}

func (p *SendUDataTxInd) Init() {
	p.code = LORAWAN_MSG_SEND_UDATA_TX_IND
}

func (p *SendUDataTxInd) String() string {
	return fmt.Sprintf("SendUDataTxInd[Status: 0x%02X, ChannelIdx: %d, DataRateIdx: %d, NumTxPackets: %d, TRXPowerLevel: %d, RFMessageAirtime: %d]", p.Status, p.ChannelIdx, p.DataRateIdx, p.NumTxPackets, p.TRXPowerLevel, p.RFMessageAirtime)
}

func (p *SendUDataTxInd) Decode(bytes []byte) error {
	p.Status = bytes[0]
	if p.Status != LORAWAN_MSG_SEND_UDATA_TX_IND_STATUS_OK && p.Status != LORAWAN_MSG_SEND_UDATA_TX_IND_STATUS_OK_ATTACHMENT {
		p.Status = LORAWAN_MSG_SEND_UDATA_TX_IND_STATUS_ERROR
		return fmt.Errorf("LORAWAN_MSG_SEND_UDATA_TX_IND_STATUS_ERROR")
	}
	if p.Status == LORAWAN_MSG_SEND_UDATA_TX_IND_STATUS_OK_ATTACHMENT {
		p.ChannelIdx = bytes[1]
		p.DataRateIdx = bytes[2]
		p.NumTxPackets = bytes[3]
		p.TRXPowerLevel = bytes[4]
		p.RFMessageAirtime = binary.LittleEndian.Uint32(bytes[5:9])
	}
	return nil
}

// LORAWAN_MSG_RECV_UDATA_IND
// LORAWAN_MSG_SEND_CDATA_REQ
// LORAWAN_MSG_SEND_CDATA_RSP
// LORAWAN_MSG_SEND_CDATA_TX_IND
// LORAWAN_MSG_RECV_CDATA_IND
// LORAWAN_MSG_RECV_ACK_IND
// LORAWAN_MSG_RECV_NO_DATA_IND
// LORAWAN_MSG_SET_RSTACK_CONFIG_REQ
// LORAWAN_MSG_SET_RSTACK_CONFIG_RSP
// LORAWAN_MSG_GET_RSTACK_CONFIG_REQ

type GetRStackConfigReq struct {
	wimodMessageImpl
}

func NewGetRStackConfigReq() *GetRStackConfigReq {
	req := &GetRStackConfigReq{}
	req.Init()
	return req
}

func (p *GetRStackConfigReq) Init() {
	p.code = LORAWAN_MSG_GET_RSTACK_CONFIG_REQ
}

func (p *GetRStackConfigReq) String() string {
	return fmt.Sprintf("GetRStackConfigReq[]")
}

func (p *GetRStackConfigReq) Encode() ([]byte, error) {
	return []byte{}, nil
}

// LORAWAN_MSG_GET_RSTACK_CONFIG_RSP

type GetRStackConfigResp struct {
	wimodMessageStatusImpl
	DefaultDataRateIdx   byte
	TXPowerLevel         byte
	AdaptativeDataRate   bool
	DutyCycleControl     bool
	ClassC               bool
	MACEvents            bool
	ExtendedHCI          bool
	AutomaticPowerSaving bool
	MaxRetransmissions   byte
	BandIdx              byte
	HeaderMACCmdCapacity byte
}

func NewGetRStackConfigResp() *GetRStackConfigResp {
	resp := &GetRStackConfigResp{}
	resp.Init()
	return resp
}

func (p *GetRStackConfigResp) Init() {
	p.code = LORAWAN_MSG_GET_RSTACK_CONFIG_RSP
}

func (p *GetRStackConfigResp) String() string {
	return fmt.Sprintf("GetRStackConfigResp[DefaultDataRateIdx: %d, TXPowerLevel: %d, AdaptativeDataRate: %t, DutyCycleControl: %t, ClassC: %t, MACEvents: %t, ExtendedHCI: %t, AutomaticPowerSaving: %t, MaxRetransmissions: %d, BandIdx: %d, HeaderMACCmdCapacity: %d]", p.DefaultDataRateIdx, p.TXPowerLevel, p.AdaptativeDataRate, p.DutyCycleControl, p.ClassC, p.MACEvents, p.ExtendedHCI, p.AutomaticPowerSaving, p.MaxRetransmissions, p.BandIdx, p.HeaderMACCmdCapacity)
}

func (p *GetRStackConfigResp) Decode(payload []byte) error {
	p.Status = payload[0]
	err := lorawanStatusCheck(p.Status)
	if err != nil {
		return err
	}
	p.DefaultDataRateIdx = payload[1]
	p.TXPowerLevel = payload[2]
	p.AdaptativeDataRate = payload[3]&0x01 == 1
	p.DutyCycleControl = (payload[3]>>1)&0x01 == 1
	p.ClassC = (payload[3]>>2)&0x01 == 1
	p.MACEvents = (payload[3]>>6)&0x01 == 1
	p.ExtendedHCI = (payload[3]>>7)&0x01 == 1
	p.AutomaticPowerSaving = payload[4]&0x01 == 1
	p.MaxRetransmissions = payload[5]
	p.BandIdx = payload[6]
	p.HeaderMACCmdCapacity = payload[7]
	return nil
}

// LORAWAN_MSG_REACTIVATE_DEVICE_REQ

type ReactivateDeviceReq struct {
	wimodMessageImpl
}

func NewReactivateDeviceReq() *ReactivateDeviceReq {
	req := &ReactivateDeviceReq{}
	req.Init()
	return req
}

func (p *ReactivateDeviceReq) Init() {
	p.code = LORAWAN_MSG_REACTIVATE_DEVICE_REQ
}

func (p *ReactivateDeviceReq) String() string {
	return fmt.Sprintf("ReactivateDeviceReq[]")
}

func (p *ReactivateDeviceReq) Encode() ([]byte, error) {
	return []byte{}, nil
}

// LORAWAN_MSG_REACTIVATE_DEVICE_RSP

type ReactivateDeviceResp struct {
	wimodMessageStatusImpl
	Address uint32
}

func NewReactivateDeviceResp() *ReactivateDeviceResp {
	resp := &ReactivateDeviceResp{}
	resp.Init()
	return resp
}

func (p *ReactivateDeviceResp) Init() {
	p.code = LORAWAN_MSG_REACTIVATE_DEVICE_RSP
}

func (p *ReactivateDeviceResp) String() string {
	return fmt.Sprintf("ReactivateDeviceResp[Address: 0x%08X]", p.Address)
}

func (p *ReactivateDeviceResp) Decode(payload []byte) error {
	p.Status = payload[0]
	err := lorawanStatusCheck(p.Status)
	if err != nil {
		return err
	}
	p.Address = binary.LittleEndian.Uint32(payload[1:5])
	return nil
}

// LORAWAN_MSG_DEACTIVATE_DEVICE_REQ

type DeactivateDeviceReq struct {
	wimodMessageImpl
}

func NewDeactivateDeviceReq() *DeactivateDeviceReq {
	req := &DeactivateDeviceReq{}
	req.Init()
	return req
}

func (p *DeactivateDeviceReq) Init() {
	p.code = LORAWAN_MSG_DEACTIVATE_DEVICE_REQ
}

func (p *DeactivateDeviceReq) String() string {
	return fmt.Sprintf("DeactivateDeviceReq[]")
}

func (p *DeactivateDeviceReq) Encode() ([]byte, error) {
	return []byte{}, nil
}

// LORAWAN_MSG_DEACTIVATE_DEVICE_RSP

type DeactivateDeviceResp struct {
	wimodMessageStatusImpl
}

func NewDeactivateDeviceResp() *DeactivateDeviceResp {
	resp := &DeactivateDeviceResp{}
	resp.Init()
	return resp
}

func (p *DeactivateDeviceResp) Init() {
	p.code = LORAWAN_MSG_DEACTIVATE_DEVICE_RSP
}

func (p *DeactivateDeviceResp) String() string {
	return fmt.Sprintf("DeactivateDeviceResp[]")
}

func (p *DeactivateDeviceResp) Decode(payload []byte) error {
	p.Status = payload[0]
	return lorawanStatusCheck(p.Status)
}

// LORAWAN_MSG_FACTORY_RESET_REQ
// LORAWAN_MSG_FACTORY_RESET_RSP
// LORAWAN_MSG_SET_DEVICE_EUI_REQ
// LORAWAN_MSG_SET_DEVICE_EUI_RSP
// LORAWAN_MSG_GET_DEVICE_EUI_REQ

type GetDeviceEUIReq struct {
	wimodMessageImpl
}

func NewGetDeviceEUIReq() *GetDeviceEUIReq {
	req := &GetDeviceEUIReq{}
	req.Init()
	return req
}

func (p *GetDeviceEUIReq) Init() {
	p.code = LORAWAN_MSG_GET_DEVICE_EUI_REQ
}

func (p *GetDeviceEUIReq) String() string {
	return fmt.Sprintf("GetDeviceEUIReq[]")
}

func (p *GetDeviceEUIReq) Encode() ([]byte, error) {
	return []byte{}, nil
}

// LORAWAN_MSG_GET_DEVICE_EUI_RSP

type GetDeviceEUIResp struct {
	wimodMessageStatusImpl
	EUI EUI
}

func NewGetDeviceEUIResp() *GetDeviceEUIResp {
	resp := &GetDeviceEUIResp{}
	resp.Init()
	return resp
}

func (p *GetDeviceEUIResp) Init() {
	p.code = LORAWAN_MSG_GET_DEVICE_EUI_RSP
}

func (p *GetDeviceEUIResp) String() string {
	return fmt.Sprintf("GetDeviceEUIResp[EUI: %v]", p.EUI)
}

func (p *GetDeviceEUIResp) Decode(payload []byte) error {
	p.Status = payload[0]
	err := lorawanStatusCheck(payload[0])
	if err != nil {
		return err
	}
	p.EUI = DecodeEUI(payload[1:])
	return nil
}

// LORAWAN_MSG_GET_NWK_STATUS_REQ

type GetNwkStatusReq struct {
	wimodMessageImpl
}

func NewGetNwkStatusReq() *GetNwkStatusReq {
	req := &GetNwkStatusReq{}
	req.Init()
	return req
}

func (p *GetNwkStatusReq) Init() {
	p.code = LORAWAN_MSG_GET_NWK_STATUS_REQ
}

func (p *GetNwkStatusReq) String() string {
	return fmt.Sprintf("GetNwkStatusReq[]")
}

func (p *GetNwkStatusReq) Encode() ([]byte, error) {
	return []byte{}, nil
}

// LORAWAN_MSG_GET_NWK_STATUS_RSP

type GetNwkStatusResp struct {
	wimodMessageStatusImpl
	NetworkStatus  byte
	Address        uint32
	DataRateIdx    byte
	PowerLevel     byte
	MaxPayloadSize byte
}

func NewGetNwkStatusResp() *GetNwkStatusResp {
	resp := &GetNwkStatusResp{}
	resp.Init()
	return resp
}

func (p *GetNwkStatusResp) Init() {
	p.code = LORAWAN_MSG_GET_NWK_STATUS_RSP
}

func (p *GetNwkStatusResp) String() string {
	return fmt.Sprintf("GetNwkStatusResp[NetworkStatus: 0x%02X, Address: 0x%08X, DataRateIdx: %d, PowerLevel: %d, MaxPayloadSize: %d]", p.NetworkStatus, p.Address, p.DataRateIdx, p.PowerLevel, p.MaxPayloadSize)
}

func (p *GetNwkStatusResp) Decode(payload []byte) error {
	p.Status = payload[0]
	err := lorawanStatusCheck(p.Status)
	if err != nil {
		return err
	}
	p.NetworkStatus = payload[1]
	if p.NetworkStatus == LORAWAN_NETWORK_STATUS_ACTIVE_ABP || p.NetworkStatus == LORAWAN_NETWORK_STATUS_ACTIVE_OTAA {
		p.Address = binary.LittleEndian.Uint32(payload[2:6])
		p.DataRateIdx = payload[6]
		p.PowerLevel = payload[7]
		p.MaxPayloadSize = payload[8]
	}
	return nil
}

// LORAWAN_MSG_SEND_MAC_CMD_REQ
// LORAWAN_MSG_SEND_MAC_CMD_RSP
// LORAWAN_MSG_RECV_MAC_CMD_IND
// LORAWAN_MSG_SET_CUSTOM_CFG_REQ
// LORAWAN_MSG_SET_CUSTOM_CFG_RSP
// LORAWAN_MSG_GET_CUSTOM_CFG_REQ
// LORAWAN_MSG_GET_CUSTOM_CFG_RSP
// LORAWAN_MSG_GET_SUPPORTED_BANDS_REQ
// LORAWAN_MSG_GET_SUPPORTED_BANDS_RSP
// LORAWAN_MSG_SET_LINKADRREQ_CONFIG_REQ
// LORAWAN_MSG_SET_LINKADRREQ_CONFIG_RSP
// LORAWAN_MSG_GET_LINKADRREQ_CONFIG_REQ
// LORAWAN_MSG_GET_LINKADRREQ_CONFIG_RSP
