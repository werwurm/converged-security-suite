//go:build !manifestcodegen
// +build !manifestcodegen

// Code generated by "menifestcodegen". DO NOT EDIT.
// To reproduce: go run github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/manifestcodegen/cmd/manifestcodegen github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/bootpolicy

package bootpolicy

import (
	"encoding/binary"
	"fmt"
	"io"
	"strings"

	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest"
	"github.com/9elements/converged-security-suite/v2/pkg/intel/metadata/manifest/common/pretty"
)

var (
	// Just to avoid errors in "import" above in case if it wasn't used below
	_ = binary.LittleEndian
	_ = (fmt.Stringer)(nil)
	_ = (io.Reader)(nil)
	_ = pretty.Header
	_ = strings.Join
	_ = manifest.StructInfo{}
)

// NewManifest returns a new instance of Manifest with
// all default values set.
func NewManifest() *Manifest {
	s := &Manifest{}
	// Recursively initializing a child structure:
	s.BPMH = *NewBPMH()
	// Recursively initializing a child structure:
	s.PMSE = *NewSignature()
	s.Rehash()
	return s
}

// Validate (recursively) checks the structure if there are any unexpected
// values. It returns an error if so.
func (s *Manifest) Validate() error {
	// Recursively validating a child structure:
	if err := s.BPMH.Validate(); err != nil {
		return fmt.Errorf("error on field 'BPMH': %w", err)
	}
	// See tag "rehashValue"
	{
		expectedValue := BPMH(s.rehashedBPMH())
		if s.BPMH != expectedValue {
			return fmt.Errorf("field 'BPMH' expects write-value '%v', but has %v", expectedValue, s.BPMH)
		}
	}
	// Recursively validating a child structure:
	if err := s.PMSE.Validate(); err != nil {
		return fmt.Errorf("error on field 'PMSE': %w", err)
	}

	return nil
}

// fieldIndexByStructID returns the position index within
// structure Manifest of the field by its StructureID
// (see document #575623, an example of StructureID value is "__KEYM__").
func (_ Manifest) fieldIndexByStructID(structID string) int {
	switch structID {
	case StructureIDBPMH:
		return 0
	case StructureIDSE:
		return 1
	case StructureIDTXT:
		return 2
	case StructureIDReserved:
		return 3
	case StructureIDPCD:
		return 4
	case StructureIDPM:
		return 5
	case StructureIDSignature:
		return 6
	}

	return -1
}

// fieldNameByIndex returns the name of the field by its position number
// within structure Manifest.
func (_ Manifest) fieldNameByIndex(fieldIndex int) string {
	switch fieldIndex {
	case 0:
		return "BPMH"
	case 1:
		return "SE"
	case 2:
		return "TXTE"
	case 3:
		return "Res"
	case 4:
		return "PCDE"
	case 5:
		return "PME"
	case 6:
		return "PMSE"
	}

	return fmt.Sprintf("invalidFieldIndex_%d", fieldIndex)
}

// ReadFrom reads the Manifest from 'r' in format defined in the document #575623.
func (s *Manifest) ReadFrom(r io.Reader) (returnN int64, returnErr error) {
	var missingFieldsByIndices = [7]bool{
		0: true,
		6: true,
	}
	defer func() {
		if returnErr != nil {
			return
		}
		for fieldIndex, v := range missingFieldsByIndices {
			if v {
				returnErr = fmt.Errorf("field '%s' is missing", s.fieldNameByIndex(fieldIndex))
				break
			}
		}
	}()
	var totalN int64
	previousFieldIndex := int(-1)
	for {
		var structInfo manifest.StructInfo
		err := binary.Read(r, binary.LittleEndian, &structInfo)
		if err == io.EOF || err == io.ErrUnexpectedEOF {
			return totalN, nil
		}
		if err != nil {
			return totalN, fmt.Errorf("unable to read structure info at %d: %w", totalN, err)
		}
		totalN += int64(binary.Size(structInfo))

		structID := structInfo.ID.String()
		fieldIndex := s.fieldIndexByStructID(structID)
		if fieldIndex < 0 {
			// TODO: report error "unknown structure ID: '"+structID+"'"
			continue
		}
		if manifest.StrictOrderCheck && fieldIndex < previousFieldIndex {
			return totalN, fmt.Errorf("invalid order of fields (%d < %d): structure '%s' is out of order", fieldIndex, previousFieldIndex, structID)
		}
		missingFieldsByIndices[fieldIndex] = false

		var n int64
		switch structID {
		case StructureIDBPMH:
			if fieldIndex == previousFieldIndex {
				return totalN, fmt.Errorf("field 'BPMH' is not a slice, but multiple elements found")
			}
			s.BPMH.SetStructInfo(structInfo)
			n, err = s.BPMH.ReadDataFrom(r)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field BPMH at %d: %w", totalN, err)
			}
		case StructureIDSE:
			var el SE
			el.SetStructInfo(structInfo)
			n, err = el.ReadDataFrom(r)
			s.SE = append(s.SE, el)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field SE at %d: %w", totalN, err)
			}
		case StructureIDTXT:
			if fieldIndex == previousFieldIndex {
				return totalN, fmt.Errorf("field 'TXTE' is not a slice, but multiple elements found")
			}
			s.TXTE = &TXT{}
			s.TXTE.SetStructInfo(structInfo)
			n, err = s.TXTE.ReadDataFrom(r)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field TXTE at %d: %w", totalN, err)
			}
		case StructureIDReserved:
			if fieldIndex == previousFieldIndex {
				return totalN, fmt.Errorf("field 'Res' is not a slice, but multiple elements found")
			}
			s.Res = &Reserved{}
			s.Res.SetStructInfo(structInfo)
			n, err = s.Res.ReadDataFrom(r)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field Res at %d: %w", totalN, err)
			}
		case StructureIDPCD:
			if fieldIndex == previousFieldIndex {
				return totalN, fmt.Errorf("field 'PCDE' is not a slice, but multiple elements found")
			}
			s.PCDE = &PCD{}
			s.PCDE.SetStructInfo(structInfo)
			n, err = s.PCDE.ReadDataFrom(r)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field PCDE at %d: %w", totalN, err)
			}
		case StructureIDPM:
			if fieldIndex == previousFieldIndex {
				return totalN, fmt.Errorf("field 'PME' is not a slice, but multiple elements found")
			}
			s.PME = &PM{}
			s.PME.SetStructInfo(structInfo)
			n, err = s.PME.ReadDataFrom(r)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field PME at %d: %w", totalN, err)
			}
		case StructureIDSignature:
			if fieldIndex == previousFieldIndex {
				return totalN, fmt.Errorf("field 'PMSE' is not a slice, but multiple elements found")
			}
			s.PMSE.SetStructInfo(structInfo)
			n, err = s.PMSE.ReadDataFrom(r)
			if err != nil {
				return totalN, fmt.Errorf("unable to read field PMSE at %d: %w", totalN, err)
			}
		default:
			return totalN, fmt.Errorf("there is no field with structure ID '%s' in Manifest", structInfo.ID)
		}
		totalN += n
		previousFieldIndex = fieldIndex
	}
}

// RehashRecursive calls Rehash (see below) recursively.
func (s *Manifest) RehashRecursive() {
	s.BPMH.Rehash()
	if s.TXTE != nil {
		s.TXTE.Rehash()
	}
	if s.Res != nil {
		s.Res.Rehash()
	}
	if s.PCDE != nil {
		s.PCDE.Rehash()
	}
	if s.PME != nil {
		s.PME.Rehash()
	}
	s.PMSE.Rehash()
	s.Rehash()
}

// Rehash sets values which are calculated automatically depending on the rest
// data. It is usually about the total size field of an element.
func (s *Manifest) Rehash() {
	s.BPMH = BPMH(s.rehashedBPMH())
}

// WriteTo writes the Manifest into 'w' in format defined in
// the document #575623.
func (s *Manifest) WriteTo(w io.Writer) (int64, error) {
	totalN := int64(0)
	s.Rehash()

	// BPMH (ManifestFieldType: element)
	{
		n, err := s.BPMH.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'BPMH': %w", err)
		}
		totalN += int64(n)
	}

	// SE (ManifestFieldType: elementList)
	{
		for idx := range s.SE {
			n, err := s.SE[idx].WriteTo(w)
			if err != nil {
				return totalN, fmt.Errorf("unable to write field 'SE[%d]': %w", idx, err)
			}
			totalN += int64(n)
		}
	}

	// TXTE (ManifestFieldType: element)
	if s.TXTE != nil {
		n, err := s.TXTE.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'TXTE': %w", err)
		}
		totalN += int64(n)
	}

	// Res (ManifestFieldType: element)
	if s.Res != nil {
		n, err := s.Res.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'Res': %w", err)
		}
		totalN += int64(n)
	}

	// PCDE (ManifestFieldType: element)
	if s.PCDE != nil {
		n, err := s.PCDE.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PCDE': %w", err)
		}
		totalN += int64(n)
	}

	// PME (ManifestFieldType: element)
	if s.PME != nil {
		n, err := s.PME.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PME': %w", err)
		}
		totalN += int64(n)
	}

	// PMSE (ManifestFieldType: element)
	{
		n, err := s.PMSE.WriteTo(w)
		if err != nil {
			return totalN, fmt.Errorf("unable to write field 'PMSE': %w", err)
		}
		totalN += int64(n)
	}

	return totalN, nil
}

// BPMHSize returns the size in bytes of the value of field BPMH
func (s *Manifest) BPMHTotalSize() uint64 {
	return s.BPMH.TotalSize()
}

// SESize returns the size in bytes of the value of field SE
func (s *Manifest) SETotalSize() uint64 {
	var size uint64
	for idx := range s.SE {
		size += s.SE[idx].TotalSize()
	}
	return size
}

// TXTESize returns the size in bytes of the value of field TXTE
func (s *Manifest) TXTETotalSize() uint64 {
	return s.TXTE.TotalSize()
}

// ResSize returns the size in bytes of the value of field Res
func (s *Manifest) ResTotalSize() uint64 {
	return s.Res.TotalSize()
}

// PCDESize returns the size in bytes of the value of field PCDE
func (s *Manifest) PCDETotalSize() uint64 {
	return s.PCDE.TotalSize()
}

// PMESize returns the size in bytes of the value of field PME
func (s *Manifest) PMETotalSize() uint64 {
	return s.PME.TotalSize()
}

// PMSESize returns the size in bytes of the value of field PMSE
func (s *Manifest) PMSETotalSize() uint64 {
	return s.PMSE.TotalSize()
}

// BPMHOffset returns the offset in bytes of field BPMH
func (s *Manifest) BPMHOffset() uint64 {
	return 0
}

// SEOffset returns the offset in bytes of field SE
func (s *Manifest) SEOffset() uint64 {
	return s.BPMHOffset() + s.BPMHTotalSize()
}

// TXTEOffset returns the offset in bytes of field TXTE
func (s *Manifest) TXTEOffset() uint64 {
	return s.SEOffset() + s.SETotalSize()
}

// ResOffset returns the offset in bytes of field Res
func (s *Manifest) ResOffset() uint64 {
	return s.TXTEOffset() + s.TXTETotalSize()
}

// PCDEOffset returns the offset in bytes of field PCDE
func (s *Manifest) PCDEOffset() uint64 {
	return s.ResOffset() + s.ResTotalSize()
}

// PMEOffset returns the offset in bytes of field PME
func (s *Manifest) PMEOffset() uint64 {
	return s.PCDEOffset() + s.PCDETotalSize()
}

// PMSEOffset returns the offset in bytes of field PMSE
func (s *Manifest) PMSEOffset() uint64 {
	return s.PMEOffset() + s.PMETotalSize()
}

// Size returns the total size of the Manifest.
func (s *Manifest) TotalSize() uint64 {
	if s == nil {
		return 0
	}

	var size uint64
	size += s.BPMHTotalSize()
	size += s.SETotalSize()
	size += s.TXTETotalSize()
	size += s.ResTotalSize()
	size += s.PCDETotalSize()
	size += s.PMETotalSize()
	size += s.PMSETotalSize()
	return size
}

// PrettyString returns the content of the structure in an easy-to-read format.
func (s *Manifest) PrettyString(depth uint, withHeader bool, opts ...pretty.Option) string {
	var lines []string
	if withHeader {
		lines = append(lines, pretty.Header(depth, "Boot Policy Manifest", s))
	}
	if s == nil {
		return strings.Join(lines, "\n")
	}
	// ManifestFieldType is element
	lines = append(lines, pretty.SubValue(depth+1, "BPMH: Header", "", &s.BPMH, opts...)...)
	// ManifestFieldType is elementList
	lines = append(lines, pretty.Header(depth+1, fmt.Sprintf("SE: Array of \"Boot Policy Manifest\" of length %d", len(s.SE)), s.SE))
	for i := 0; i < len(s.SE); i++ {
		lines = append(lines, fmt.Sprintf("%sitem #%d: ", strings.Repeat("  ", int(depth+2)), i)+strings.TrimSpace(s.SE[i].PrettyString(depth+2, true)))
	}
	if depth < 1 {
		lines = append(lines, "")
	}
	// ManifestFieldType is element
	lines = append(lines, pretty.SubValue(depth+1, "TXTE", "", s.TXTE, opts...)...)
	// ManifestFieldType is element
	lines = append(lines, pretty.SubValue(depth+1, "Res", "", s.Res, opts...)...)
	// ManifestFieldType is element
	lines = append(lines, pretty.SubValue(depth+1, "PCDE: Platform Config Data", "", s.PCDE, opts...)...)
	// ManifestFieldType is element
	lines = append(lines, pretty.SubValue(depth+1, "PME: Platform Manufacturer", "", s.PME, opts...)...)
	// ManifestFieldType is element
	lines = append(lines, pretty.SubValue(depth+1, "PMSE: Signature", "", &s.PMSE, opts...)...)
	if depth < 2 {
		lines = append(lines, "")
	}
	return strings.Join(lines, "\n")
}
