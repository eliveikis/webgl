// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package webgl

import (
	"github.com/gopherjs/gopherwasm/js"
)

type ContextAttributes struct {
	// If Alpha is true, the drawing buffer has an alpha channel for
	// the purposes of performing OpenGL destination alpha operations
	// and compositing with the page.
	Alpha bool

	// If Depth is true, the drawing buffer has a depth buffer of at least 16 bits.
	Depth bool

	// If Stencil is true, the drawing buffer has a stencil buffer of at least 8 bits.
	Stencil bool

	// If Antialias is true and the implementation supports antialiasing
	// the drawing buffer will perform antialiasing using its choice of
	// technique (multisample/supersample) and quality.
	Antialias bool

	// If PremultipliedAlpha is true the page compositor will assume the
	// drawing buffer contains colors with premultiplied alpha.
	// This flag is ignored if the alpha flag is false.
	PremultipliedAlpha bool

	// If the value is true the buffers will not be cleared and will preserve
	// their values until cleared or overwritten by the author.
	PreserveDrawingBuffer bool
}

// Returns a copy of the default WebGL context attributes.
func DefaultAttributes() *ContextAttributes {
	return &ContextAttributes{true, true, false, true, true, false}
}

type Context struct {
	*js.Value
	ARRAY_BUFFER                                 int `js:"ARRAY_BUFFER"`
	ARRAY_BUFFER_BINDING                         int `js:"ARRAY_BUFFER_BINDING"`
	ATTACHED_SHADERS                             int `js:"ATTACHED_SHADERS"`
	BACK                                         int `js:"BACK"`
	BLEND                                        int `js:"BLEND"`
	BLEND_COLOR                                  int `js:"BLEND_COLOR"`
	BLEND_DST_ALPHA                              int `js:"BLEND_DST_ALPHA"`
	BLEND_DST_RGB                                int `js:"BLEND_DST_RGB"`
	BLEND_EQUATION                               int `js:"BLEND_EQUATION"`
	BLEND_EQUATION_ALPHA                         int `js:"BLEND_EQUATION_ALPHA"`
	BLEND_EQUATION_RGB                           int `js:"BLEND_EQUATION_RGB"`
	BLEND_SRC_ALPHA                              int `js:"BLEND_SRC_ALPHA"`
	BLEND_SRC_RGB                                int `js:"BLEND_SRC_RGB"`
	BLUE_BITS                                    int `js:"BLUE_BITS"`
	BOOL                                         int `js:"BOOL"`
	BOOL_VEC2                                    int `js:"BOOL_VEC2"`
	BOOL_VEC3                                    int `js:"BOOL_VEC3"`
	BOOL_VEC4                                    int `js:"BOOL_VEC4"`
	BROWSER_DEFAULT_WEBGL                        int `js:"BROWSER_DEFAULT_WEBGL"`
	BUFFER_SIZE                                  int `js:"BUFFER_SIZE"`
	BUFFER_USAGE                                 int `js:"BUFFER_USAGE"`
	BYTE                                         int `js:"BYTE"`
	CCW                                          int `js:"CCW"`
	CLAMP_TO_EDGE                                int `js:"CLAMP_TO_EDGE"`
	COLOR_ATTACHMENT0                            int `js:"COLOR_ATTACHMENT0"`
	COLOR_BUFFER_BIT                             int `js:"COLOR_BUFFER_BIT"`
	COLOR_CLEAR_VALUE                            int `js:"COLOR_CLEAR_VALUE"`
	COLOR_WRITEMASK                              int `js:"COLOR_WRITEMASK"`
	COMPILE_STATUS                               int `js:"COMPILE_STATUS"`
	COMPRESSED_TEXTURE_FORMATS                   int `js:"COMPRESSED_TEXTURE_FORMATS"`
	CONSTANT_ALPHA                               int `js:"CONSTANT_ALPHA"`
	CONSTANT_COLOR                               int `js:"CONSTANT_COLOR"`
	CONTEXT_LOST_WEBGL                           int `js:"CONTEXT_LOST_WEBGL"`
	CULL_FACE                                    int `js:"CULL_FACE"`
	CULL_FACE_MODE                               int `js:"CULL_FACE_MODE"`
	CURRENT_PROGRAM                              int `js:"CURRENT_PROGRAM"`
	CURRENT_VERTEX_ATTRIB                        int `js:"CURRENT_VERTEX_ATTRIB"`
	CW                                           int `js:"CW"`
	DECR                                         int `js:"DECR"`
	DECR_WRAP                                    int `js:"DECR_WRAP"`
	DELETE_STATUS                                int `js:"DELETE_STATUS"`
	DEPTH_ATTACHMENT                             int `js:"DEPTH_ATTACHMENT"`
	DEPTH_BITS                                   int `js:"DEPTH_BITS"`
	DEPTH_BUFFER_BIT                             int `js:"DEPTH_BUFFER_BIT"`
	DEPTH_CLEAR_VALUE                            int `js:"DEPTH_CLEAR_VALUE"`
	DEPTH_COMPONENT                              int `js:"DEPTH_COMPONENT"`
	DEPTH_COMPONENT16                            int `js:"DEPTH_COMPONENT16"`
	DEPTH_FUNC                                   int `js:"DEPTH_FUNC"`
	DEPTH_RANGE                                  int `js:"DEPTH_RANGE"`
	DEPTH_STENCIL                                int `js:"DEPTH_STENCIL"`
	DEPTH_STENCIL_ATTACHMENT                     int `js:"DEPTH_STENCIL_ATTACHMENT"`
	DEPTH_TEST                                   int `js:"DEPTH_TEST"`
	DEPTH_WRITEMASK                              int `js:"DEPTH_WRITEMASK"`
	DITHER                                       int `js:"DITHER"`
	DONT_CARE                                    int `js:"DONT_CARE"`
	DST_ALPHA                                    int `js:"DST_ALPHA"`
	DST_COLOR                                    int `js:"DST_COLOR"`
	DYNAMIC_DRAW                                 int `js:"DYNAMIC_DRAW"`
	ELEMENT_ARRAY_BUFFER                         int `js:"ELEMENT_ARRAY_BUFFER"`
	ELEMENT_ARRAY_BUFFER_BINDING                 int `js:"ELEMENT_ARRAY_BUFFER_BINDING"`
	EQUAL                                        int `js:"EQUAL"`
	FASTEST                                      int `js:"FASTEST"`
	FLOAT                                        int `js:"FLOAT"`
	FLOAT_MAT2                                   int `js:"FLOAT_MAT2"`
	FLOAT_MAT3                                   int `js:"FLOAT_MAT3"`
	FLOAT_MAT4                                   int `js:"FLOAT_MAT4"`
	FLOAT_VEC2                                   int `js:"FLOAT_VEC2"`
	FLOAT_VEC3                                   int `js:"FLOAT_VEC3"`
	FLOAT_VEC4                                   int `js:"FLOAT_VEC4"`
	FRAGMENT_SHADER                              int `js:"FRAGMENT_SHADER"`
	FRAMEBUFFER                                  int `js:"FRAMEBUFFER"`
	FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           int `js:"FRAMEBUFFER_ATTACHMENT_OBJECT_NAME"`
	FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           int `js:"FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE"`
	FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE int `js:"FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE"`
	FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         int `js:"FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL"`
	FRAMEBUFFER_BINDING                          int `js:"FRAMEBUFFER_BINDING"`
	FRAMEBUFFER_COMPLETE                         int `js:"FRAMEBUFFER_COMPLETE"`
	FRAMEBUFFER_INCOMPLETE_ATTACHMENT            int `js:"FRAMEBUFFER_INCOMPLETE_ATTACHMENT"`
	FRAMEBUFFER_INCOMPLETE_DIMENSIONS            int `js:"FRAMEBUFFER_INCOMPLETE_DIMENSIONS"`
	FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    int `js:"FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"`
	FRAMEBUFFER_UNSUPPORTED                      int `js:"FRAMEBUFFER_UNSUPPORTED"`
	FRONT                                        int `js:"FRONT"`
	FRONT_AND_BACK                               int `js:"FRONT_AND_BACK"`
	FRONT_FACE                                   int `js:"FRONT_FACE"`
	FUNC_ADD                                     int `js:"FUNC_ADD"`
	FUNC_REVERSE_SUBTRACT                        int `js:"FUNC_REVERSE_SUBTRACT"`
	FUNC_SUBTRACT                                int `js:"FUNC_SUBTRACT"`
	GENERATE_MIPMAP_HINT                         int `js:"GENERATE_MIPMAP_HINT"`
	GEQUAL                                       int `js:"GEQUAL"`
	GREATER                                      int `js:"GREATER"`
	GREEN_BITS                                   int `js:"GREEN_BITS"`
	HIGH_FLOAT                                   int `js:"HIGH_FLOAT"`
	HIGH_INT                                     int `js:"HIGH_INT"`
	INCR                                         int `js:"INCR"`
	INCR_WRAP                                    int `js:"INCR_WRAP"`
	INFO_LOG_LENGTH                              int `js:"INFO_LOG_LENGTH"`
	INT                                          int `js:"INT"`
	INT_VEC2                                     int `js:"INT_VEC2"`
	INT_VEC3                                     int `js:"INT_VEC3"`
	INT_VEC4                                     int `js:"INT_VEC4"`
	INVALID_ENUM                                 int `js:"INVALID_ENUM"`
	INVALID_FRAMEBUFFER_OPERATION                int `js:"INVALID_FRAMEBUFFER_OPERATION"`
	INVALID_OPERATION                            int `js:"INVALID_OPERATION"`
	INVALID_VALUE                                int `js:"INVALID_VALUE"`
	INVERT                                       int `js:"INVERT"`
	KEEP                                         int `js:"KEEP"`
	LEQUAL                                       int `js:"LEQUAL"`
	LESS                                         int `js:"LESS"`
	LINEAR                                       int `js:"LINEAR"`
	LINEAR_MIPMAP_LINEAR                         int `js:"LINEAR_MIPMAP_LINEAR"`
	LINEAR_MIPMAP_NEAREST                        int `js:"LINEAR_MIPMAP_NEAREST"`
	LINES                                        int `js:"LINES"`
	LINE_LOOP                                    int `js:"LINE_LOOP"`
	LINE_STRIP                                   int `js:"LINE_STRIP"`
	LINE_WIDTH                                   int `js:"LINE_WIDTH"`
	LINK_STATUS                                  int `js:"LINK_STATUS"`
	LOW_FLOAT                                    int `js:"LOW_FLOAT"`
	LOW_INT                                      int `js:"LOW_INT"`
	LUMINANCE                                    int `js:"LUMINANCE"`
	LUMINANCE_ALPHA                              int `js:"LUMINANCE_ALPHA"`
	MAX_COMBINED_TEXTURE_IMAGE_UNITS             int `js:"MAX_COMBINED_TEXTURE_IMAGE_UNITS"`
	MAX_CUBE_MAP_TEXTURE_SIZE                    int `js:"MAX_CUBE_MAP_TEXTURE_SIZE"`
	MAX_FRAGMENT_UNIFORM_VECTORS                 int `js:"MAX_FRAGMENT_UNIFORM_VECTORS"`
	MAX_RENDERBUFFER_SIZE                        int `js:"MAX_RENDERBUFFER_SIZE"`
	MAX_TEXTURE_IMAGE_UNITS                      int `js:"MAX_TEXTURE_IMAGE_UNITS"`
	MAX_TEXTURE_SIZE                             int `js:"MAX_TEXTURE_SIZE"`
	MAX_VARYING_VECTORS                          int `js:"MAX_VARYING_VECTORS"`
	MAX_VERTEX_ATTRIBS                           int `js:"MAX_VERTEX_ATTRIBS"`
	MAX_VERTEX_TEXTURE_IMAGE_UNITS               int `js:"MAX_VERTEX_TEXTURE_IMAGE_UNITS"`
	MAX_VERTEX_UNIFORM_VECTORS                   int `js:"MAX_VERTEX_UNIFORM_VECTORS"`
	MAX_VIEWPORT_DIMS                            int `js:"MAX_VIEWPORT_DIMS"`
	MEDIUM_FLOAT                                 int `js:"MEDIUM_FLOAT"`
	MEDIUM_INT                                   int `js:"MEDIUM_INT"`
	MIRRORED_REPEAT                              int `js:"MIRRORED_REPEAT"`
	NEAREST                                      int `js:"NEAREST"`
	NEAREST_MIPMAP_LINEAR                        int `js:"NEAREST_MIPMAP_LINEAR"`
	NEAREST_MIPMAP_NEAREST                       int `js:"NEAREST_MIPMAP_NEAREST"`
	NEVER                                        int `js:"NEVER"`
	NICEST                                       int `js:"NICEST"`
	NONE                                         int `js:"NONE"`
	NOTEQUAL                                     int `js:"NOTEQUAL"`
	NO_ERROR                                     int `js:"NO_ERROR"`
	NUM_COMPRESSED_TEXTURE_FORMATS               int `js:"NUM_COMPRESSED_TEXTURE_FORMATS"`
	ONE                                          int `js:"ONE"`
	ONE_MINUS_CONSTANT_ALPHA                     int `js:"ONE_MINUS_CONSTANT_ALPHA"`
	ONE_MINUS_CONSTANT_COLOR                     int `js:"ONE_MINUS_CONSTANT_COLOR"`
	ONE_MINUS_DST_ALPHA                          int `js:"ONE_MINUS_DST_ALPHA"`
	ONE_MINUS_DST_COLOR                          int `js:"ONE_MINUS_DST_COLOR"`
	ONE_MINUS_SRC_ALPHA                          int `js:"ONE_MINUS_SRC_ALPHA"`
	ONE_MINUS_SRC_COLOR                          int `js:"ONE_MINUS_SRC_COLOR"`
	OUT_OF_MEMORY                                int `js:"OUT_OF_MEMORY"`
	PACK_ALIGNMENT                               int `js:"PACK_ALIGNMENT"`
	POINTS                                       int `js:"POINTS"`
	POLYGON_OFFSET_FACTOR                        int `js:"POLYGON_OFFSET_FACTOR"`
	POLYGON_OFFSET_FILL                          int `js:"POLYGON_OFFSET_FILL"`
	POLYGON_OFFSET_UNITS                         int `js:"POLYGON_OFFSET_UNITS"`
	RED_BITS                                     int `js:"RED_BITS"`
	RENDERBUFFER                                 int `js:"RENDERBUFFER"`
	RENDERBUFFER_ALPHA_SIZE                      int `js:"RENDERBUFFER_ALPHA_SIZE"`
	RENDERBUFFER_BINDING                         int `js:"RENDERBUFFER_BINDING"`
	RENDERBUFFER_BLUE_SIZE                       int `js:"RENDERBUFFER_BLUE_SIZE"`
	RENDERBUFFER_DEPTH_SIZE                      int `js:"RENDERBUFFER_DEPTH_SIZE"`
	RENDERBUFFER_GREEN_SIZE                      int `js:"RENDERBUFFER_GREEN_SIZE"`
	RENDERBUFFER_HEIGHT                          int `js:"RENDERBUFFER_HEIGHT"`
	RENDERBUFFER_INTERNAL_FORMAT                 int `js:"RENDERBUFFER_INTERNAL_FORMAT"`
	RENDERBUFFER_RED_SIZE                        int `js:"RENDERBUFFER_RED_SIZE"`
	RENDERBUFFER_STENCIL_SIZE                    int `js:"RENDERBUFFER_STENCIL_SIZE"`
	RENDERBUFFER_WIDTH                           int `js:"RENDERBUFFER_WIDTH"`
	RENDERER                                     int `js:"RENDERER"`
	REPEAT                                       int `js:"REPEAT"`
	REPLACE                                      int `js:"REPLACE"`
	RGB                                          int `js:"RGB"`
	RGB5_A1                                      int `js:"RGB5_A1"`
	RGB565                                       int `js:"RGB565"`
	RGBA                                         int `js:"RGBA"`
	RGBA4                                        int `js:"RGBA4"`
	SAMPLER_2D                                   int `js:"SAMPLER_2D"`
	SAMPLER_CUBE                                 int `js:"SAMPLER_CUBE"`
	SAMPLES                                      int `js:"SAMPLES"`
	SAMPLE_ALPHA_TO_COVERAGE                     int `js:"SAMPLE_ALPHA_TO_COVERAGE"`
	SAMPLE_BUFFERS                               int `js:"SAMPLE_BUFFERS"`
	SAMPLE_COVERAGE                              int `js:"SAMPLE_COVERAGE"`
	SAMPLE_COVERAGE_INVERT                       int `js:"SAMPLE_COVERAGE_INVERT"`
	SAMPLE_COVERAGE_VALUE                        int `js:"SAMPLE_COVERAGE_VALUE"`
	SCISSOR_BOX                                  int `js:"SCISSOR_BOX"`
	SCISSOR_TEST                                 int `js:"SCISSOR_TEST"`
	SHADER_COMPILER                              int `js:"SHADER_COMPILER"`
	SHADER_SOURCE_LENGTH                         int `js:"SHADER_SOURCE_LENGTH"`
	SHADER_TYPE                                  int `js:"SHADER_TYPE"`
	SHADING_LANGUAGE_VERSION                     int `js:"SHADING_LANGUAGE_VERSION"`
	SHORT                                        int `js:"SHORT"`
	SRC_ALPHA                                    int `js:"SRC_ALPHA"`
	SRC_ALPHA_SATURATE                           int `js:"SRC_ALPHA_SATURATE"`
	SRC_COLOR                                    int `js:"SRC_COLOR"`
	STATIC_DRAW                                  int `js:"STATIC_DRAW"`
	STENCIL_ATTACHMENT                           int `js:"STENCIL_ATTACHMENT"`
	STENCIL_BACK_FAIL                            int `js:"STENCIL_BACK_FAIL"`
	STENCIL_BACK_FUNC                            int `js:"STENCIL_BACK_FUNC"`
	STENCIL_BACK_PASS_DEPTH_FAIL                 int `js:"STENCIL_BACK_PASS_DEPTH_FAIL"`
	STENCIL_BACK_PASS_DEPTH_PASS                 int `js:"STENCIL_BACK_PASS_DEPTH_PASS"`
	STENCIL_BACK_REF                             int `js:"STENCIL_BACK_REF"`
	STENCIL_BACK_VALUE_MASK                      int `js:"STENCIL_BACK_VALUE_MASK"`
	STENCIL_BACK_WRITEMASK                       int `js:"STENCIL_BACK_WRITEMASK"`
	STENCIL_BITS                                 int `js:"STENCIL_BITS"`
	STENCIL_BUFFER_BIT                           int `js:"STENCIL_BUFFER_BIT"`
	STENCIL_CLEAR_VALUE                          int `js:"STENCIL_CLEAR_VALUE"`
	STENCIL_FAIL                                 int `js:"STENCIL_FAIL"`
	STENCIL_FUNC                                 int `js:"STENCIL_FUNC"`
	STENCIL_INDEX                                int `js:"STENCIL_INDEX"`
	STENCIL_INDEX8                               int `js:"STENCIL_INDEX8"`
	STENCIL_PASS_DEPTH_FAIL                      int `js:"STENCIL_PASS_DEPTH_FAIL"`
	STENCIL_PASS_DEPTH_PASS                      int `js:"STENCIL_PASS_DEPTH_PASS"`
	STENCIL_REF                                  int `js:"STENCIL_REF"`
	STENCIL_TEST                                 int `js:"STENCIL_TEST"`
	STENCIL_VALUE_MASK                           int `js:"STENCIL_VALUE_MASK"`
	STENCIL_WRITEMASK                            int `js:"STENCIL_WRITEMASK"`
	STREAM_DRAW                                  int `js:"STREAM_DRAW"`
	SUBPIXEL_BITS                                int `js:"SUBPIXEL_BITS"`
	TEXTURE                                      int `js:"TEXTURE"`
	TEXTURE0                                     int `js:"TEXTURE0"`
	TEXTURE1                                     int `js:"TEXTURE1"`
	TEXTURE2                                     int `js:"TEXTURE2"`
	TEXTURE3                                     int `js:"TEXTURE3"`
	TEXTURE4                                     int `js:"TEXTURE4"`
	TEXTURE5                                     int `js:"TEXTURE5"`
	TEXTURE6                                     int `js:"TEXTURE6"`
	TEXTURE7                                     int `js:"TEXTURE7"`
	TEXTURE8                                     int `js:"TEXTURE8"`
	TEXTURE9                                     int `js:"TEXTURE9"`
	TEXTURE10                                    int `js:"TEXTURE10"`
	TEXTURE11                                    int `js:"TEXTURE11"`
	TEXTURE12                                    int `js:"TEXTURE12"`
	TEXTURE13                                    int `js:"TEXTURE13"`
	TEXTURE14                                    int `js:"TEXTURE14"`
	TEXTURE15                                    int `js:"TEXTURE15"`
	TEXTURE16                                    int `js:"TEXTURE16"`
	TEXTURE17                                    int `js:"TEXTURE17"`
	TEXTURE18                                    int `js:"TEXTURE18"`
	TEXTURE19                                    int `js:"TEXTURE19"`
	TEXTURE20                                    int `js:"TEXTURE20"`
	TEXTURE21                                    int `js:"TEXTURE21"`
	TEXTURE22                                    int `js:"TEXTURE22"`
	TEXTURE23                                    int `js:"TEXTURE23"`
	TEXTURE24                                    int `js:"TEXTURE24"`
	TEXTURE25                                    int `js:"TEXTURE25"`
	TEXTURE26                                    int `js:"TEXTURE26"`
	TEXTURE27                                    int `js:"TEXTURE27"`
	TEXTURE28                                    int `js:"TEXTURE28"`
	TEXTURE29                                    int `js:"TEXTURE29"`
	TEXTURE30                                    int `js:"TEXTURE30"`
	TEXTURE31                                    int `js:"TEXTURE31"`
	TEXTURE_2D                                   int `js:"TEXTURE_2D"`
	TEXTURE_BINDING_2D                           int `js:"TEXTURE_BINDING_2D"`
	TEXTURE_BINDING_CUBE_MAP                     int `js:"TEXTURE_BINDING_CUBE_MAP"`
	TEXTURE_CUBE_MAP                             int `js:"TEXTURE_CUBE_MAP"`
	TEXTURE_CUBE_MAP_NEGATIVE_X                  int `js:"TEXTURE_CUBE_MAP_NEGATIVE_X"`
	TEXTURE_CUBE_MAP_NEGATIVE_Y                  int `js:"TEXTURE_CUBE_MAP_NEGATIVE_Y"`
	TEXTURE_CUBE_MAP_NEGATIVE_Z                  int `js:"TEXTURE_CUBE_MAP_NEGATIVE_Z"`
	TEXTURE_CUBE_MAP_POSITIVE_X                  int `js:"TEXTURE_CUBE_MAP_POSITIVE_X"`
	TEXTURE_CUBE_MAP_POSITIVE_Y                  int `js:"TEXTURE_CUBE_MAP_POSITIVE_Y"`
	TEXTURE_CUBE_MAP_POSITIVE_Z                  int `js:"TEXTURE_CUBE_MAP_POSITIVE_Z"`
	TEXTURE_MAG_FILTER                           int `js:"TEXTURE_MAG_FILTER"`
	TEXTURE_MIN_FILTER                           int `js:"TEXTURE_MIN_FILTER"`
	TEXTURE_WRAP_S                               int `js:"TEXTURE_WRAP_S"`
	TEXTURE_WRAP_T                               int `js:"TEXTURE_WRAP_T"`
	TRIANGLES                                    int `js:"TRIANGLES"`
	TRIANGLE_FAN                                 int `js:"TRIANGLE_FAN"`
	TRIANGLE_STRIP                               int `js:"TRIANGLE_STRIP"`
	UNPACK_ALIGNMENT                             int `js:"UNPACK_ALIGNMENT"`
	UNPACK_COLORSPACE_CONVERSION_WEBGL           int `js:"UNPACK_COLORSPACE_CONVERSION_WEBGL"`
	UNPACK_FLIP_Y_WEBGL                          int `js:"UNPACK_FLIP_Y_WEBGL"`
	UNPACK_PREMULTIPLY_ALPHA_WEBGL               int `js:"UNPACK_PREMULTIPLY_ALPHA_WEBGL"`
	UNSIGNED_BYTE                                int `js:"UNSIGNED_BYTE"`
	UNSIGNED_INT                                 int `js:"UNSIGNED_INT"`
	UNSIGNED_SHORT                               int `js:"UNSIGNED_SHORT"`
	UNSIGNED_SHORT_4_4_4_4                       int `js:"UNSIGNED_SHORT_4_4_4_4"`
	UNSIGNED_SHORT_5_5_5_1                       int `js:"UNSIGNED_SHORT_5_5_5_1"`
	UNSIGNED_SHORT_5_6_5                         int `js:"UNSIGNED_SHORT_5_6_5"`
	VALIDATE_STATUS                              int `js:"VALIDATE_STATUS"`
	VENDOR                                       int `js:"VENDOR"`
	VERSION                                      int `js:"VERSION"`
	VERTEX_ATTRIB_ARRAY_BUFFER_BINDING           int `js:"VERTEX_ATTRIB_ARRAY_BUFFER_BINDING"`
	VERTEX_ATTRIB_ARRAY_ENABLED                  int `js:"VERTEX_ATTRIB_ARRAY_ENABLED"`
	VERTEX_ATTRIB_ARRAY_NORMALIZED               int `js:"VERTEX_ATTRIB_ARRAY_NORMALIZED"`
	VERTEX_ATTRIB_ARRAY_POINTER                  int `js:"VERTEX_ATTRIB_ARRAY_POINTER"`
	VERTEX_ATTRIB_ARRAY_SIZE                     int `js:"VERTEX_ATTRIB_ARRAY_SIZE"`
	VERTEX_ATTRIB_ARRAY_STRIDE                   int `js:"VERTEX_ATTRIB_ARRAY_STRIDE"`
	VERTEX_ATTRIB_ARRAY_TYPE                     int `js:"VERTEX_ATTRIB_ARRAY_TYPE"`
	VERTEX_SHADER                                int `js:"VERTEX_SHADER"`
	VIEWPORT                                     int `js:"VIEWPORT"`
	ZERO                                         int `js:"ZERO"`
}
