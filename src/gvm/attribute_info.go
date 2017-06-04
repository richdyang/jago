package gvm

/*
attribute_info {
    u2 attribute_name_index;
    u4 attribute_length;
    u1 info[attribute_length];
}
*/
type AttributeInfo interface {
	ReadInfo(reader *ClassReader)
}

/*
Code_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 max_stack;
    u2 max_locals;
    u4 code_length;
    u1 code[code_length];
    u2 exception_table_length;
    {   u2 start_pc;
        u2 end_pc;
        u2 handler_pc;
        u2 catch_type;
    } exception_table[exception_table_length];
    u2 attributes_count;
    attribute_info attributes[attributes_count];
}
*/
type CodeAttribute struct {
	attributeNameIndex   uint16
	attributeLength      uint32
	maxStack             uint16
	maxLocals            uint16
	codeLength           uint32
	code                 []uint8               //u4 code_length
	exceptionTableLength uint16
	exceptionTable       []ExceptionTableEntry //u2 exception_table_length
	attributesCount      uint16
	attributes           []AttributeInfo  //u2 attributes_count
}

func (this *CodeAttribute) ReadInfo(reader *ClassReader) {
	this.maxStack = reader.ReadUint16()
	this.maxLocals = reader.ReadUint16()
	codeLength := reader.ReadUint32()
	this.codeLength = codeLength
	this.code = reader.ReadBytes(codeLength)
	exceptionTableLength := reader.ReadUint16()
	this.exceptionTableLength = exceptionTableLength
	this.exceptionTable = make([]ExceptionTableEntry, exceptionTableLength)
	for i := uint16(0); i < exceptionTableLength; i++ {
		exceptionTable := ExceptionTableEntry{}
		exceptionTable.ReadInfo(reader)
		this.exceptionTable[i] = exceptionTable
	}
	attributesCount := reader.ReadUint16()
	this.attributesCount = attributesCount
	this.attributes = make([]AttributeInfo, attributesCount)
	for i := uint16(0); i < attributesCount; i++ {
		reader.readAttribute(&this.attributes[i])
	}
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (this *ExceptionTableEntry) ReadInfo(reader *ClassReader) {
	this.startPc = reader.ReadUint16()
	this.endPc = reader.ReadUint16()
	this.handlerPc = reader.ReadUint16()
	this.catchType = reader.ReadUint16()
}

/*
LineNumberTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 line_number_table_length;
    {   u2 start_pc;
        u2 line_number;
    } line_number_table[line_number_table_length];
}
*/
type LineNumberTableAttribute struct {
	attributeNameIndex    uint16
	attributeLength       uint32
	lineNumberTableLength uint16
	lineNumberTable       []LineNumberTableEntry
}

type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (this *LineNumberTableAttribute) ReadInfo(reader *ClassReader) {
	lineNumberTableLength := reader.ReadUint16()
	this.lineNumberTableLength = lineNumberTableLength
	this.lineNumberTable = make([]LineNumberTableEntry, lineNumberTableLength)
	for i := uint16(0); i < lineNumberTableLength; i++ {
		this.lineNumberTable[i] = LineNumberTableEntry{reader.ReadUint16(), reader.ReadUint16()}
	}
}

/*
SourceFile_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 sourcefile_index;
}
 */

type SourceFileAttribue struct {
	attributeNameIndex uint16
	attributeLength    uint32
	sourceFileIndex uint16
}

func (this *SourceFileAttribue) ReadInfo(reader *ClassReader) {
	this.sourceFileIndex = reader.ReadUint16()
}

/*
LocalVariableTable_attribute {
    u2 attribute_name_index;
    u4 attribute_length;
    u2 local_variable_table_length;
    {   u2 start_pc;
        u2 length;
        u2 name_index;
        u2 descriptor_index;
        u2 index;
    } local_variable_table[local_variable_table_length];
}
 */
type LocalVariableTableAttribute struct {
	attributeNameIndex       uint16
	attributeLength          uint32
	localVariableTableLength uint16
	localVariableTable       []LocalVariableTableEntry
}

type LocalVariableTableEntry struct {
	startPc             uint16
	length              uint16
	nameIndex           uint16
	descriptorIndex     uint16
	index               uint16
}

func (this *LocalVariableTableAttribute) ReadInfo(reader *ClassReader) {
	localVariableTableLength := reader.ReadUint16()
	this.localVariableTableLength = localVariableTableLength
	this.localVariableTable = make([]LocalVariableTableEntry, localVariableTableLength)
	for i := uint16(0); i < localVariableTableLength; i++ {
		this.localVariableTable[i] = LocalVariableTableEntry{
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16(),
			reader.ReadUint16()}
	}
}

