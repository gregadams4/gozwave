package commands
type DeviceResetLocallyNotification struct {
    *report
    data []byte
}

func NewDeviceResetLocallyNotification(data []byte) *DeviceResetLocallyNotification {
    if len(data) < 0 {
				//may want to change this to just return nil
				for i := len(data); i < 0; i++ {
            data = append(data, 0x00)
        }
    }

    return &DeviceResetLocallyNotification{
        data: data,
    }
}

