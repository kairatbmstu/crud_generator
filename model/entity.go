package model

/**
 Field types Permalink to "Field types"

JHipster supports many field types. This support depends on your database backend, so we use Java types to describe them: a Java String will be stored differently in Oracle or Cassandra, and it is one of JHipster’s strengths to generate the correct database access code for you.

    String: A Java String. Its default size depends on the underlying backend (if you use JPA, it’s 255 by default), but you can change it using the validation rules (putting a max size of 1024, for example).
    Integer: A Java Integer.
    Long: A Java Long.
    Float: A Java Float.
    Double: A Java Double.
    BigDecimal: A java.math.BigDecimal object, used when you want exact mathematic calculations (often used for financial operations).
    LocalDate: A java.time.LocalDate object, used to correctly manage dates in Java.
    Instant: A java.time.Instant object, used to represent a timestamp, an instantaneous point on the time-line.
    ZonedDateTime: A java.time.ZonedDateTime object, used to represent a local date-time in a given timezone (typically a calendar appointment). Note that time zones are neither supported by the REST nor by the persistence layers so you should most probably use Instant instead.
    Duration: A java.time.Duration object, used to represent an amount of time.
    UUID: A java.util.UUID.
    Boolean: A Java Boolean.
    Enumeration: A Java Enumeration object. When this type is selected, the sub-generator will ask you what values you want in your enumeration, and it will create a specific enum class to store them.
    Blob: A Blob object, used to store some binary data. When this type is selected, the sub-generator will ask you if you want to store generic binary data, an image object, or a CLOB (long text). Images will be handled specifically on the Angular side, so they can be displayed to the end-user.

*/
const (
	FieldType_int     = "Integer"
	FieldType_string  = "String"
	FieldType_uuid    = "UUID"
	FieldType_boolean = "Boolean"
)

type FieldType string

const (
	PaginationType_InfiniteScroll = "infinite-scroll"
	PaginationType_Pagination     = "pagination"
)

type PaginationType string

type Field struct {
	Name             string
	ColumnName       string
	JsonName         string
	Type             FieldType
	IsRequired       bool
	PaginatationType PaginationType
}

type Entity struct {
	Name   string
	Fields []Field
}

type RelationshipType string

const (
	RelationshipType_OneToOne   = "onetoone"
	RelationshipType_OneToMany  = "onetomany"
	RelationshipType_ManyToOne  = "manytoone"
	RelationshipType_ManyToMany = "manytomany"
)

type Relationship struct {
	EntityOne      Entity
	EntityOneField Field
	EntityTwo      Entity
	EntityTwoField Field
	RelationType   RelationshipType
}

type Paginate struct {
	EntityName     string
	PaginationType string
}

type File struct {
	Name string
}

type Model struct {
	Entities      []Entity
	Relationships []Relationship
}
