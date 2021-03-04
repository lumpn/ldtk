// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    ldtkJSON, err := UnmarshalLdtkJSON(bytes)
//    bytes, err = ldtkJSON.Marshal()

package main

import "encoding/json"

func UnmarshalLdtkJSON(data []byte) (LdtkJSON, error) {
	var r LdtkJSON
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *LdtkJSON) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

// This file is a JSON schema of files created by LDtk level editor (https://ldtk.io).
//
// This is the root of any Project JSON file. It contains:  - the project settings, - an
// array of levels, - and a definition object (that can probably be safely ignored for most
// users).
type LdtkJSON struct {
	BackupLimit         int64       `json:"backupLimit"`        // Number of backup files to keep, if the `backupOnSave` is TRUE
	BackupOnSave        bool        `json:"backupOnSave"`       // If TRUE, an extra copy of the project will be created in a sub folder, when saving.
	BgColor             string      `json:"bgColor"`            // Project background color
	DefaultGridSize     int64       `json:"defaultGridSize"`    // Default grid size for new layers
	DefaultLevelBgColor string      `json:"defaultLevelBgColor"`// Default background color of levels
	DefaultLevelHeight  int64       `json:"defaultLevelHeight"` // Default new level height
	DefaultLevelWidth   int64       `json:"defaultLevelWidth"`  // Default new level width
	DefaultPivotX       float64     `json:"defaultPivotX"`      // Default X pivot (0 to 1) for new entities
	DefaultPivotY       float64     `json:"defaultPivotY"`      // Default Y pivot (0 to 1) for new entities
	Defs                Definitions `json:"defs"`               // A structure containing all the definitions of this project
	ExportPNG           bool        `json:"exportPng"`          // If TRUE, all layers in all levels will also be exported as PNG along with the project; file (default is FALSE)
	ExportTiled         bool        `json:"exportTiled"`        // If TRUE, a Tiled compatible file will also be generated along with the LDtk JSON file; (default is FALSE)
	ExternalLevels      bool        `json:"externalLevels"`     // If TRUE, one file will be saved for the project (incl. all its definitions) and one file; in a sub-folder for each level.
	Flags               []Flag      `json:"flags"`              // An array containing various advanced flags (ie. options or other states). Possible; values: `DiscardPreCsvIntGrid`, `IgnoreBackupSuggest`
	JSONVersion         string      `json:"jsonVersion"`        // File format version
	Levels              []Level     `json:"levels"`             // All levels. The order of this array is only relevant in `LinearHorizontal` and; `linearVertical` world layouts (see `worldLayout` value). Otherwise, you should refer to; the `worldX`,`worldY` coordinates of each Level.
	MinifyJSON          bool        `json:"minifyJson"`         // If TRUE, the Json is partially minified (no indentation, nor line breaks, default is; FALSE)
	NextUid             int64       `json:"nextUid"`            // Next Unique integer ID available
	PNGFilePattern      *string     `json:"pngFilePattern"`     // File naming pattern for exported PNGs
	WorldGridHeight     int64       `json:"worldGridHeight"`    // Height of the world grid in pixels.
	WorldGridWidth      int64       `json:"worldGridWidth"`     // Width of the world grid in pixels.
	WorldLayout         WorldLayout `json:"worldLayout"`        // An enum that describes how levels are organized in this project (ie. linearly or in a 2D; space). Possible values: `Free`, `GridVania`, `LinearHorizontal`, `LinearVertical`
}

// A structure containing all the definitions of this project
//
// If you're writing your own LDtk importer, you should probably just ignore *most* stuff in
// the `defs` section, as it contains data that are mostly important to the editor. To keep
// you away from the `defs` section and avoid some unnecessary JSON parsing, important data
// from definitions is often duplicated in fields prefixed with a double underscore (eg.
// `__identifier` or `__type`).  The 2 only definition types you might need here are
// **Tilesets** and **Enums**.
type Definitions struct {
	Entities      []EntityDefinition  `json:"entities"`     // All entities, including their custom fields
	Enums         []EnumDefinition    `json:"enums"`        
	ExternalEnums []EnumDefinition    `json:"externalEnums"`// Note: external enums are exactly the same as `enums`, except they have a `relPath` to; point to an external source file.
	Layers        []LayerDefinition   `json:"layers"`       
	LevelFields   []FieldDefinition   `json:"levelFields"`  // An array containing all custom fields available to all levels.
	Tilesets      []TilesetDefinition `json:"tilesets"`     
}

type EntityDefinition struct {
	Color           string            `json:"color"`          // Base entity color
	FieldDefs       []FieldDefinition `json:"fieldDefs"`      // Array of field definitions
	FillOpacity     float64           `json:"fillOpacity"`    
	Height          int64             `json:"height"`         // Pixel height
	Hollow          bool              `json:"hollow"`         
	Identifier      string            `json:"identifier"`     // Unique String identifier
	KeepAspectRatio bool              `json:"keepAspectRatio"`// Only applies to entities resizable on both X/Y. If TRUE, the entity instance width/height; will keep the same aspect ratio as the definition.
	LimitBehavior   LimitBehavior     `json:"limitBehavior"`  // Possible values: `DiscardOldOnes`, `PreventAdding`, `MoveLastOne`
	LimitScope      LimitScope        `json:"limitScope"`     // If TRUE, the maxCount is a "per world" limit, if FALSE, it's a "per level". Possible; values: `PerLayer`, `PerLevel`, `PerWorld`
	LineOpacity     float64           `json:"lineOpacity"`    
	MaxCount        int64             `json:"maxCount"`       // Max instances count
	PivotX          float64           `json:"pivotX"`         // Pivot X coordinate (from 0 to 1.0)
	PivotY          float64           `json:"pivotY"`         // Pivot Y coordinate (from 0 to 1.0)
	RenderMode      RenderMode        `json:"renderMode"`     // Possible values: `Rectangle`, `Ellipse`, `Tile`, `Cross`
	ResizableX      bool              `json:"resizableX"`     // If TRUE, the entity instances will be resizable horizontally
	ResizableY      bool              `json:"resizableY"`     // If TRUE, the entity instances will be resizable vertically
	ShowName        bool              `json:"showName"`       // Display entity name in editor
	Tags            []string          `json:"tags"`           // An array of strings that classifies this entity
	TileID          *int64            `json:"tileId"`         // Tile ID used for optional tile display
	TileRenderMode  TileRenderMode    `json:"tileRenderMode"` // Possible values: `Cover`, `FitInside`, `Repeat`, `Stretch`
	TilesetID       *int64            `json:"tilesetId"`      // Tileset ID used for optional tile display
	Uid             int64             `json:"uid"`            // Unique Int identifier
	Width           int64             `json:"width"`          // Pixel width
}

// This section is mostly only intended for the LDtk editor app itself. You can safely
// ignore it.
type FieldDefinition struct {
	Type                string            `json:"__type"`             // Human readable value type (eg. `Int`, `Float`, `Point`, etc.). If the field is an array,; this field will look like `Array<...>` (eg. `Array<Int>`, `Array<Point>` etc.)
	AcceptFileTypes     []string          `json:"acceptFileTypes"`    // Optional list of accepted file extensions for FilePath value type. Includes the dot:; `.ext`
	ArrayMaxLength      *int64            `json:"arrayMaxLength"`     // Array max length
	ArrayMinLength      *int64            `json:"arrayMinLength"`     // Array min length
	CanBeNull           bool              `json:"canBeNull"`          // TRUE if the value can be null. For arrays, TRUE means it can contain null values; (exception: array of Points can't have null values).
	DefaultOverride     interface{}       `json:"defaultOverride"`    // Default value if selected value is null or invalid.
	EditorAlwaysShow    bool              `json:"editorAlwaysShow"`   
	EditorCutLongValues bool              `json:"editorCutLongValues"`
	EditorDisplayMode   EditorDisplayMode `json:"editorDisplayMode"`  // Possible values: `Hidden`, `ValueOnly`, `NameAndValue`, `EntityTile`, `PointStar`,; `PointPath`, `RadiusPx`, `RadiusGrid`
	EditorDisplayPos    EditorDisplayPos  `json:"editorDisplayPos"`   // Possible values: `Above`, `Center`, `Beneath`
	Identifier          string            `json:"identifier"`         // Unique String identifier
	IsArray             bool              `json:"isArray"`            // TRUE if the value is an array of multiple values
	Max                 *float64          `json:"max"`                // Max limit for value, if applicable
	Min                 *float64          `json:"min"`                // Min limit for value, if applicable
	Regex               *string           `json:"regex"`              // Optional regular expression that needs to be matched to accept values. Expected format:; `/some_reg_ex/g`, with optional "i" flag.
	TextLangageMode     *TextLangageMode  `json:"textLangageMode"`    // Possible values: &lt;`null`&gt;, `LangPython`, `LangRuby`, `LangJS`, `LangLua`, `LangC`,; `LangHaxe`, `LangMarkdown`, `LangJson`, `LangXml`
	FieldDefinitionType interface{}       `json:"type"`               // Internal type enum
	Uid                 int64             `json:"uid"`                // Unique Intidentifier
}

type EnumDefinition struct {
	ExternalFileChecksum *string               `json:"externalFileChecksum"`
	ExternalRelPath      *string               `json:"externalRelPath"`     // Relative path to the external file providing this Enum
	IconTilesetUid       *int64                `json:"iconTilesetUid"`      // Tileset UID if provided
	Identifier           string                `json:"identifier"`          // Unique String identifier
	Uid                  int64                 `json:"uid"`                 // Unique Int identifier
	Values               []EnumValueDefinition `json:"values"`              // All possible enum values, with their optional Tile infos.
}

type EnumValueDefinition struct {
	TileSrcRect []int64 `json:"__tileSrcRect"`// An array of 4 Int values that refers to the tile in the tileset image: `[ x, y, width,; height ]`
	ID          string  `json:"id"`           // Enum value
	TileID      *int64  `json:"tileId"`       // The optional ID of the tile
}

type LayerDefinition struct {
	Type                  string                   `json:"__type"`               // Type of the layer (*IntGrid, Entities, Tiles or AutoLayer*)
	AutoRuleGroups        []map[string]interface{} `json:"autoRuleGroups"`       // Contains all the auto-layer rule definitions.
	AutoSourceLayerDefUid *int64                   `json:"autoSourceLayerDefUid"`
	AutoTilesetDefUid     *int64                   `json:"autoTilesetDefUid"`    // Reference to the Tileset UID being used by this auto-layer rules
	DisplayOpacity        float64                  `json:"displayOpacity"`       // Opacity of the layer (0 to 1.0)
	ExcludedTags          []string                 `json:"excludedTags"`         // An array of tags to forbid some Entities in this layer
	GridSize              int64                    `json:"gridSize"`             // Width and height of the grid in pixels
	Identifier            string                   `json:"identifier"`           // Unique String identifier
	IntGridValues         []IntGridValueDefinition `json:"intGridValues"`        // An array that defines extra optional info for each IntGrid value. The array is sorted; using value (ascending).
	PxOffsetX             int64                    `json:"pxOffsetX"`            // X offset of the layer, in pixels (IMPORTANT: this should be added to the `LayerInstance`; optional offset)
	PxOffsetY             int64                    `json:"pxOffsetY"`            // Y offset of the layer, in pixels (IMPORTANT: this should be added to the `LayerInstance`; optional offset)
	RequiredTags          []string                 `json:"requiredTags"`         // An array of tags to filter Entities that can be added to this layer
	TilePivotX            float64                  `json:"tilePivotX"`           // If the tiles are smaller or larger than the layer grid, the pivot value will be used to; position the tile relatively its grid cell.
	TilePivotY            float64                  `json:"tilePivotY"`           // If the tiles are smaller or larger than the layer grid, the pivot value will be used to; position the tile relatively its grid cell.
	TilesetDefUid         *int64                   `json:"tilesetDefUid"`        // Reference to the Tileset UID being used by this Tile layer
	LayerDefinitionType   Type                     `json:"type"`                 // Type of the layer as Haxe Enum Possible values: `IntGrid`, `Entities`, `Tiles`,; `AutoLayer`
	Uid                   int64                    `json:"uid"`                  // Unique Int identifier
}

// IntGrid value definition
type IntGridValueDefinition struct {
	Color      string  `json:"color"`     
	Identifier *string `json:"identifier"`// Unique String identifier
	Value      int64   `json:"value"`     // The IntGrid value itself
}

// The `Tileset` definition is the most important part among project definitions. It
// contains some extra informations about each integrated tileset. If you only had to parse
// one definition section, that would be the one.
type TilesetDefinition struct {
	CachedPixelData map[string]interface{}   `json:"cachedPixelData"`// The following data is used internally for various optimizations. It's always synced with; source image changes.
	Identifier      string                   `json:"identifier"`     // Unique String identifier
	Padding         int64                    `json:"padding"`        // Distance in pixels from image borders
	PxHei           int64                    `json:"pxHei"`          // Image height in pixels
	PxWid           int64                    `json:"pxWid"`          // Image width in pixels
	RelPath         string                   `json:"relPath"`        // Path to the source file, relative to the current project JSON file
	SavedSelections []map[string]interface{} `json:"savedSelections"`// Array of group of tiles selections, only meant to be used in the editor
	Spacing         int64                    `json:"spacing"`        // Space in pixels between all tiles
	TileGridSize    int64                    `json:"tileGridSize"`   
	Uid             int64                    `json:"uid"`            // Unique Intidentifier
}

// This section contains all the level data. It can be found in 2 distinct forms, depending
// on Project current settings:  - If "*Separate level files*" is **disabled** (default):
// full level data is *embedded* inside the main Project JSON file, - If "*Separate level
// files*" is **enabled**: level data is stored in *separate* standalone `.ldtkl` files (one
// per level). In this case, the main Project JSON file will still contain most level data,
// except heavy sections, like the `layerInstances` array (which will be null). The
// `externalRelPath` string points to the `ldtkl` file.  A `ldtkl` file is just a JSON file
// containing exactly what is described below.
type Level struct {
	BgColor         string                   `json:"__bgColor"`      // Background color of the level (same as `bgColor`, except the default value is; automatically used here if its value is `null`)
	BgPos           *LevelBackgroundPosition `json:"__bgPos"`        // Position informations of the background image, if there is one.
	Neighbours      []NeighbourLevel         `json:"__neighbours"`   // An array listing all other levels touching this one on the world map. In "linear" world; layouts, this array is populated with previous/next levels in array, and `dir` depends on; the linear horizontal/vertical layout.
	LevelBgColor    *string                  `json:"bgColor"`        // Background color of the level. If `null`, the project `defaultLevelBgColor` should be; used.
	BgPivotX        float64                  `json:"bgPivotX"`       // Background image X pivot (0-1)
	BgPivotY        float64                  `json:"bgPivotY"`       // Background image Y pivot (0-1)
	LevelBgPos      *BgPos                   `json:"bgPos"`          // An enum defining the way the background image (if any) is positioned on the level. See; `__bgPos` for resulting position info. Possible values: &lt;`null`&gt;, `Unscaled`,; `Contain`, `Cover`, `CoverDirty`
	BgRelPath       *string                  `json:"bgRelPath"`      // The *optional* relative path to the level background image.
	ExternalRelPath *string                  `json:"externalRelPath"`// This value is not null if the project option "*Save levels separately*" is enabled. In; this case, this **relative** path points to the level Json file.
	FieldInstances  []FieldInstance          `json:"fieldInstances"` // An array containing this level custom field values.
	Identifier      string                   `json:"identifier"`     // Unique String identifier
	LayerInstances  []LayerInstance          `json:"layerInstances"` // An array containing all Layer instances. **IMPORTANT**: if the project option "*Save; levels separately*" is enabled, this field will be `null`.<br/>  This array is **sorted; in display order**: the 1st layer is the top-most and the last is behind.
	PxHei           int64                    `json:"pxHei"`          // Height of the level in pixels
	PxWid           int64                    `json:"pxWid"`          // Width of the level in pixels
	Uid             int64                    `json:"uid"`            // Unique Int identifier
	WorldX          int64                    `json:"worldX"`         // World X coordinate in pixels
	WorldY          int64                    `json:"worldY"`         // World Y coordinate in pixels
}

// Level background image position info
type LevelBackgroundPosition struct {
	CropRect  []float64 `json:"cropRect"` // An array of 4 float values describing the cropped sub-rectangle of the displayed; background image. This cropping happens when original is larger than the level bounds.; Array format: `[ cropX, cropY, cropWidth, cropHeight ]`
	Scale     []float64 `json:"scale"`    // An array containing the `[scaleX,scaleY]` values of the **cropped** background image,; depending on `bgPos` option.
	TopLeftPx []int64   `json:"topLeftPx"`// An array containing the `[x,y]` pixel coordinates of the top-left corner of the; **cropped** background image, depending on `bgPos` option.
}

type FieldInstance struct {
	Identifier       string        `json:"__identifier"`    // Field definition identifier
	Type             string        `json:"__type"`          // Type of the field, such as `Int`, `Float`, `Enum(my_enum_name)`, `Bool`, etc.
	Value            interface{}   `json:"__value"`         // Actual value of the field instance. The value type may vary, depending on `__type`; (Integer, Boolean, String etc.)<br/>  It can also be an `Array` of those same types.
	DefUid           int64         `json:"defUid"`          // Reference of the **Field definition** UID
	RealEditorValues []interface{} `json:"realEditorValues"`// Editor internal raw values
}

type LayerInstance struct {
	CHei               int64                  `json:"__cHei"`            // Grid-based height
	CWid               int64                  `json:"__cWid"`            // Grid-based width
	GridSize           int64                  `json:"__gridSize"`        // Grid size
	Identifier         string                 `json:"__identifier"`      // Layer definition identifier
	Opacity            float64                `json:"__opacity"`         // Layer opacity as Float [0-1]
	PxTotalOffsetX     int64                  `json:"__pxTotalOffsetX"`  // Total layer X pixel offset, including both instance and definition offsets.
	PxTotalOffsetY     int64                  `json:"__pxTotalOffsetY"`  // Total layer Y pixel offset, including both instance and definition offsets.
	TilesetDefUid      *int64                 `json:"__tilesetDefUid"`   // The definition UID of corresponding Tileset, if any.
	TilesetRelPath     *string                `json:"__tilesetRelPath"`  // The relative path to corresponding Tileset, if any.
	Type               string                 `json:"__type"`            // Layer type (possible values: IntGrid, Entities, Tiles or AutoLayer)
	AutoLayerTiles     []TileInstance         `json:"autoLayerTiles"`    // An array containing all tiles generated by Auto-layer rules. The array is already sorted; in display order (ie. 1st tile is beneath 2nd, which is beneath 3rd etc.).<br/><br/>; Note: if multiple tiles are stacked in the same cell as the result of different rules,; all tiles behind opaque ones will be discarded.
	EntityInstances    []EntityInstance       `json:"entityInstances"`   
	GridTiles          []TileInstance         `json:"gridTiles"`         
	IntGrid            []IntGridValueInstance `json:"intGrid,omitempty"` // **WARNING**: this deprecated value will be *removed* completely on version 0.9.0+; Replaced by: `intGridCsv`
	IntGridCSV         []int64                `json:"intGridCsv"`        // A list of all values in the IntGrid layer, stored from left to right, and top to bottom; (ie. first row from left to right, followed by second row, etc). `0` means "empty cell"; and IntGrid values start at 1. This array size is `__cWid` x `__cHei` cells.
	LayerDefUid        int64                  `json:"layerDefUid"`       // Reference the Layer definition UID
	LevelID            int64                  `json:"levelId"`           // Reference to the UID of the level containing this layer instance
	OverrideTilesetUid *int64                 `json:"overrideTilesetUid"`// This layer can use another tileset by overriding the tileset UID here.
	PxOffsetX          int64                  `json:"pxOffsetX"`         // X offset in pixels to render this layer, usually 0 (IMPORTANT: this should be added to; the `LayerDef` optional offset, see `__pxTotalOffsetX`)
	PxOffsetY          int64                  `json:"pxOffsetY"`         // Y offset in pixels to render this layer, usually 0 (IMPORTANT: this should be added to; the `LayerDef` optional offset, see `__pxTotalOffsetY`)
	Seed               int64                  `json:"seed"`              // Random seed used for Auto-Layers rendering
	Visible            bool                   `json:"visible"`           // Layer instance visibility
}

// This structure represents a single tile from a given Tileset.
type TileInstance struct {
	D   []int64 `json:"d"`  // Internal data used by the editor.<br/>  For auto-layer tiles: `[ruleId, coordId]`.<br/>; For tile-layer tiles: `[coordId]`.
	F   int64   `json:"f"`  // "Flip bits", a 2-bits integer to represent the mirror transformations of the tile.<br/>; - Bit 0 = X flip<br/>   - Bit 1 = Y flip<br/>   Examples: f=0 (no flip), f=1 (X flip; only), f=2 (Y flip only), f=3 (both flips)
	Px  []int64 `json:"px"` // Pixel coordinates of the tile in the **layer** (`[x,y]` format). Don't forget optional; layer offsets, if they exist!
	Src []int64 `json:"src"`// Pixel coordinates of the tile in the **tileset** (`[x,y]` format)
	T   int64   `json:"t"`  // The *Tile ID* in the corresponding tileset.
}

type EntityInstance struct {
	Grid           []int64             `json:"__grid"`        // Grid-based coordinates (`[x,y]` format)
	Identifier     string              `json:"__identifier"`  // Entity definition identifier
	Pivot          []float64           `json:"__pivot"`       // Pivot coordinates  (`[x,y]` format, values are from 0 to 1) of the Entity
	Tile           *EntityInstanceTile `json:"__tile"`        // Optional Tile used to display this entity (it could either be the default Entity tile, or; some tile provided by a field value, like an Enum).
	DefUid         int64               `json:"defUid"`        // Reference of the **Entity definition** UID
	FieldInstances []FieldInstance     `json:"fieldInstances"`// An array of all custom fields and their values.
	Height         int64               `json:"height"`        // Entity height in pixels. For non-resizable entities, it will be the same as Entity; definition.
	Px             []int64             `json:"px"`            // Pixel coordinates (`[x,y]` format) in current level coordinate space. Don't forget; optional layer offsets, if they exist!
	Width          int64               `json:"width"`         // Entity width in pixels. For non-resizable entities, it will be the same as Entity; definition.
}

// Tile data in an Entity instance
type EntityInstanceTile struct {
	SrcRect    []int64 `json:"srcRect"`   // An array of 4 Int values that refers to the tile in the tileset image: `[ x, y, width,; height ]`
	TilesetUid int64   `json:"tilesetUid"`// Tileset ID
}

// IntGrid value instance
type IntGridValueInstance struct {
	CoordID int64 `json:"coordId"`// Coordinate ID in the layer grid
	V       int64 `json:"v"`      // IntGrid value
}

// Nearby level info
type NeighbourLevel struct {
	Dir      string `json:"dir"`     // A single lowercase character tipping on the level location (`n`orth, `s`outh, `w`est,; `e`ast).
	LevelUid int64  `json:"levelUid"`
}

// Possible values: `Hidden`, `ValueOnly`, `NameAndValue`, `EntityTile`, `PointStar`,
// `PointPath`, `RadiusPx`, `RadiusGrid`
type EditorDisplayMode string
const (
	EntityTile EditorDisplayMode = "EntityTile"
	Hidden EditorDisplayMode = "Hidden"
	NameAndValue EditorDisplayMode = "NameAndValue"
	PointPath EditorDisplayMode = "PointPath"
	PointStar EditorDisplayMode = "PointStar"
	RadiusGrid EditorDisplayMode = "RadiusGrid"
	RadiusPx EditorDisplayMode = "RadiusPx"
	ValueOnly EditorDisplayMode = "ValueOnly"
)

// Possible values: `Above`, `Center`, `Beneath`
type EditorDisplayPos string
const (
	Above EditorDisplayPos = "Above"
	Beneath EditorDisplayPos = "Beneath"
	Center EditorDisplayPos = "Center"
)

type TextLangageMode string
const (
	LangC TextLangageMode = "LangC"
	LangHaxe TextLangageMode = "LangHaxe"
	LangJS TextLangageMode = "LangJS"
	LangJSON TextLangageMode = "LangJson"
	LangLua TextLangageMode = "LangLua"
	LangMarkdown TextLangageMode = "LangMarkdown"
	LangPython TextLangageMode = "LangPython"
	LangRuby TextLangageMode = "LangRuby"
	LangXML TextLangageMode = "LangXml"
)

// Possible values: `DiscardOldOnes`, `PreventAdding`, `MoveLastOne`
type LimitBehavior string
const (
	DiscardOldOnes LimitBehavior = "DiscardOldOnes"
	MoveLastOne LimitBehavior = "MoveLastOne"
	PreventAdding LimitBehavior = "PreventAdding"
)

// If TRUE, the maxCount is a "per world" limit, if FALSE, it's a "per level". Possible
// values: `PerLayer`, `PerLevel`, `PerWorld`
type LimitScope string
const (
	PerLayer LimitScope = "PerLayer"
	PerLevel LimitScope = "PerLevel"
	PerWorld LimitScope = "PerWorld"
)

// Possible values: `Rectangle`, `Ellipse`, `Tile`, `Cross`
type RenderMode string
const (
	Cross RenderMode = "Cross"
	Ellipse RenderMode = "Ellipse"
	Rectangle RenderMode = "Rectangle"
	Tile RenderMode = "Tile"
)

// Possible values: `Cover`, `FitInside`, `Repeat`, `Stretch`
type TileRenderMode string
const (
	FitInside TileRenderMode = "FitInside"
	Repeat TileRenderMode = "Repeat"
	Stretch TileRenderMode = "Stretch"
	TileRenderModeCover TileRenderMode = "Cover"
)

// Type of the layer as Haxe Enum Possible values: `IntGrid`, `Entities`, `Tiles`,
// `AutoLayer`
type Type string
const (
	AutoLayer Type = "AutoLayer"
	Entities Type = "Entities"
	IntGrid Type = "IntGrid"
	Tiles Type = "Tiles"
)

type Flag string
const (
	DiscardPreCSVIntGrid Flag = "DiscardPreCsvIntGrid"
	IgnoreBackupSuggest Flag = "IgnoreBackupSuggest"
)

type BgPos string
const (
	BgPosCover BgPos = "Cover"
	Contain BgPos = "Contain"
	CoverDirty BgPos = "CoverDirty"
	Unscaled BgPos = "Unscaled"
)

// An enum that describes how levels are organized in this project (ie. linearly or in a 2D
// space). Possible values: `Free`, `GridVania`, `LinearHorizontal`, `LinearVertical`
type WorldLayout string
const (
	Free WorldLayout = "Free"
	GridVania WorldLayout = "GridVania"
	LinearHorizontal WorldLayout = "LinearHorizontal"
	LinearVertical WorldLayout = "LinearVertical"
)
