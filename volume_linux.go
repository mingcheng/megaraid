package diskutil

import (
	"errors"
	"strings"
)

func (v *VirtualDriveStat) parseLine(line string) error {
	if strings.HasPrefix(line, keyVdVirtualDrive) {
		parts := strings.SplitN(line, "(", 2)
		if len(parts) != 2 {
			return errors.New("format illegal: " + line)
		}
		virtualDrive, err := parseFiled(parts[0], keyVdVirtualDrive, typeInt)
		if err != nil {
			return err
		}
		v.VirtualDrive = virtualDrive.(int)
	} else if strings.HasPrefix(line, keyVdName) || strings.HasPrefix(line, keyVdVirtualDriveType) {
		name, err := parseFiled(line, keyVdName, typeString)
		if err != nil {
			return err
		}
		v.Name = name.(string)
	} else if strings.HasPrefix(line, keyVdSize) {
		size, err := parseFiled(line, keyVdSize, typeString)
		if err != nil {
			return err
		}
		v.Size = size.(string)
	} else if strings.HasPrefix(line, keyRAIDLevel) {
		size, err := parseFiled(line, keyRAIDLevel, typeString)
		if err != nil {
			return err
		}
		v.RAIDLevel = size.(string)
	} else if strings.HasPrefix(line, keyVdState) {
		state, err := parseFiled(line, keyVdState, typeString)
		if err != nil {
			return err
		}
		v.State = state.(string)
	} else if strings.HasPrefix(line, keyVdNumberOfDrives) {
		numberOfDrives, err := parseFiled(line, keyVdNumberOfDrives, typeInt)
		if err != nil {
			return err
		}
		v.NumberOfDrives = numberOfDrives.(int)
	} else if strings.HasPrefix(line, keyVdEncryptiontype) {
		encryptiontype, err := parseFiled(line, keyVdEncryptiontype, typeString)
		if err != nil {
			return err
		}
		v.Encryptiontype = encryptiontype.(string)
	}
	return nil
}
