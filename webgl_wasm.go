// Copyright 2014 Joseph Hager. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build wasm

package webgl

import (
	"errors"

	"github.com/gopherjs/gopherwasm/js"
)

var null js.Value = js.Null()

// NewContext takes an HTML5 canvas object and optional context attributes.
// If an error is returned it means you won't have access to WebGL
// functionality.
func NewContext(canvas js.Value, ca *ContextAttributes) (*Context, error) {
	if js.Global().Get("WebGLRenderingContext") == js.Undefined() {
		return nil, errors.New("Your browser doesn't appear to support webgl.")
	}

	if ca == nil {
		ca = DefaultAttributes()
	}

	attrs := map[string]bool{
		"alpha":                 ca.Alpha,
		"depth":                 ca.Depth,
		"stencil":               ca.Stencil,
		"antialias":             ca.Antialias,
		"premultipliedAlpha":    ca.PremultipliedAlpha,
		"preserveDrawingBuffer": ca.PreserveDrawingBuffer,
	}
	gl := canvas.Call("getContext", "webgl", attrs)
	if gl == null {
		gl = canvas.Call("getContext", "experimental-webgl", attrs)
		if gl == null {
			return nil, errors.New("Creating a webgl context has failed.")
		}
	}
	ctx := new(Context)
	ctx.Value = &gl
	return ctx, nil
}

// Returns the context attributes active on the context. These values might
// be different than what was requested on context creation if the
// browser's implementation doesn't support a feature.
func (c *Context) GetContextAttributes() ContextAttributes {
	ca := c.Call("getContextAttributes")
	return ContextAttributes{
		ca.Get("alpha").Bool(),
		ca.Get("depth").Bool(),
		ca.Get("stencil").Bool(),
		ca.Get("antialias").Bool(),
		ca.Get("premultipliedAlpha").Bool(),
		ca.Get("preservedDrawingBuffer").Bool(),
	}
}

// Specifies the active texture unit.
func (c *Context) ActiveTexture(texture int) {
	c.Call("activeTexture", texture)
}

// Attaches a WebGLShader object to a WebGLProgram object.
func (c *Context) AttachShader(program js.Value, shader js.Value) {
	c.Call("attachShader", program, shader)
}

// Binds a generic vertex index to a user-defined attribute variable.
func (c *Context) BindAttribLocation(program js.Value, index int, name string) {
	c.Call("bindAttribLocation", program, index, name)
}

// Associates a buffer with a buffer target.
func (c *Context) BindBuffer(target int, buffer js.Value) {
	c.Call("bindBuffer", target, buffer)
}

// Associates a WebGLFramebuffer object with the FRAMEBUFFER bind target.
func (c *Context) BindFramebuffer(target int, framebuffer js.Value) {
	c.Call("bindFramebuffer", target, framebuffer)
}

// Binds a WebGLRenderbuffer object to be used for rendering.
func (c *Context) BindRenderbuffer(target int, renderbuffer js.Value) {
	c.Call("bindRenderbuffer", target, renderbuffer)
}

// Binds a named texture object to a target.
func (c *Context) BindTexture(target int, texture js.Value) {
	c.Call("bindTexture", target, texture)
}

// The GL_BLEND_COLOR may be used to calculate the source and destination blending factors.
func (c *Context) BlendColor(r, g, b, a float64) {
	c.Call("blendColor", r, g, b, a)
}

// Sets the equation used to blend RGB and Alpha values of an incoming source
// fragment with a destination values as stored in the fragment's frame buffer.
func (c *Context) BlendEquation(mode int) {
	c.Call("blendEquation", mode)
}

// Controls the blending of an incoming source fragment's R, G, B, and A values
// with a destination R, G, B, and A values as stored in the fragment's WebGLFramebuffer.
func (c *Context) BlendEquationSeparate(modeRGB, modeAlpha int) {
	c.Call("blendEquationSeparate", modeRGB, modeAlpha)
}

// Sets the blending factors used to combine source and destination pixels.
func (c *Context) BlendFunc(sfactor, dfactor int) {
	c.Call("blendFunc", sfactor, dfactor)
}

// Sets the weighting factors that are used by blendEquationSeparate.
func (c *Context) BlendFuncSeparate(srcRGB, dstRGB, srcAlpha, dstAlpha int) {
	c.Call("blendFuncSeparate", srcRGB, dstRGB, srcAlpha, dstAlpha)
}

// Creates a buffer in memory and initializes it with array data.
// If no array is provided, the contents of the buffer is initialized to 0.
func (c *Context) BufferData(target int, data interface{}, usage int) {
	c.Call("bufferData", target, data, usage)
}

// Used to modify or update some or all of a data store for a bound buffer object.
func (c *Context) BufferSubData(target int, offset int, data interface{}) {
	c.Call("bufferSubData", target, offset, data)
}

// Returns whether the currently bound WebGLFramebuffer is complete.
// If not complete, returns the reason why.
func (c *Context) CheckFramebufferStatus(target int) int {
	return c.Call("checkFramebufferStatus", target).Int()
}

// Sets all pixels in a specific buffer to the same value.
func (c *Context) Clear(flags int) {
	c.Call("clear", flags)
}

// Specifies color values to use by the clear method to clear the color buffer.
func (c *Context) ClearColor(r, g, b, a float32) {
	c.Call("clearColor", r, g, b, a)
}

// Clears the depth buffer to a specific value.
func (c *Context) ClearDepth(depth float64) {
	c.Call("clearDepth", depth)
}

func (c *Context) ClearStencil(s int) {
	c.Call("clearStencil", s)
}

// Lets you set whether individual colors can be written when
// drawing or rendering to a framebuffer.
func (c *Context) ColorMask(r, g, b, a bool) {
	c.Call("colorMask", r, g, b, a)
}

// Compiles the GLSL shader source into binary data used by the WebGLProgram object.
func (c *Context) CompileShader(shader js.Value) {
	c.Call("compileShader", shader)
}

// Copies a rectangle of pixels from the current WebGLFramebuffer into a texture image.
func (c *Context) CopyTexImage2D(target, level, internal, x, y, w, h, border int) {
	c.Call("copyTexImage2D", target, level, internal, x, y, w, h, border)
}

// Replaces a portion of an existing 2D texture image with data from the current framebuffer.
func (c *Context) CopyTexSubImage2D(target, level, xoffset, yoffset, x, y, w, h int) {
	c.Call("copyTexSubImage2D", target, level, xoffset, yoffset, x, y, w, h)
}

// Creates and initializes a WebGLBuffer.
func (c *Context) CreateBuffer() js.Value {
	return c.Call("createBuffer")
}

// Returns a WebGLFramebuffer object.
func (c *Context) CreateFramebuffer() js.Value {
	return c.Call("createFramebuffer")
}

// Creates an empty WebGLProgram object to which vector and fragment
// WebGLShader objects can be bound.
func (c *Context) CreateProgram() js.Value {
	return c.Call("createProgram")
}

// Creates and returns a WebGLRenderbuffer object.
func (c *Context) CreateRenderbuffer() js.Value {
	return c.Call("createRenderbuffer")
}

// Returns an empty vertex or fragment shader object based on the type specified.
func (c *Context) CreateShader(typ int) js.Value {
	return c.Call("createShader", typ)
}

// Used to generate a WebGLTexture object to which images can be bound.
func (c *Context) CreateTexture() js.Value {
	return c.Call("createTexture")
}

// Sets whether or not front, back, or both facing facets are able to be culled.
func (c *Context) CullFace(mode int) {
	c.Call("cullFace", mode)
}

// Delete a specific buffer.
func (c *Context) DeleteBuffer(buffer js.Value) {
	c.Call("deleteBuffer", buffer)
}

// Deletes a specific WebGLFramebuffer object. If you delete the
// currently bound framebuffer, the default framebuffer will be bound.
// Deleting a framebuffer detaches all of its attachments.
func (c *Context) DeleteFramebuffer(framebuffer js.Value) {
	c.Call("deleteFramebuffer", framebuffer)
}

// Flags a specific WebGLProgram object for deletion if currently active.
// It will be deleted when it is no longer being used.
// Any shader objects associated with the program will be detached.
// They will be deleted if they were already flagged for deletion.
func (c *Context) DeleteProgram(program js.Value) {
	c.Call("deleteProgram", program)
}

// Deletes the specified renderbuffer object. If the renderbuffer is
// currently bound, it will become unbound. If the renderbuffer is
// attached to the currently bound framebuffer, it is detached.
func (c *Context) DeleteRenderbuffer(renderbuffer js.Value) {
	c.Call("deleteRenderbuffer", renderbuffer)
}

// Deletes a specific shader object.
func (c *Context) DeleteShader(shader js.Value) {
	c.Call("deleteShader", shader)
}

// Deletes a specific texture object.
func (c *Context) DeleteTexture(texture js.Value) {
	c.Call("deleteTexture", texture)
}

// Sets a function to use to compare incoming pixel depth to the
// current depth buffer value.
func (c *Context) DepthFunc(fun int) {
	c.Call("depthFunc", fun)
}

// Sets whether or not you can write to the depth buffer.
func (c *Context) DepthMask(flag bool) {
	c.Call("depthMask", flag)
}

// Sets the depth range for normalized coordinates to canvas or viewport depth coordinates.
func (c *Context) DepthRange(zNear, zFar float64) {
	c.Call("depthRange", zNear, zFar)
}

// Detach a shader object from a program object.
func (c *Context) DetachShader(program, shader js.Value) {
	c.Call("detachShader", program, shader)
}

// Turns off specific WebGL capabilities for this context.
func (c *Context) Disable(cap int) {
	c.Call("disable", cap)
}

// Turns off a vertex attribute array at a specific index position.
func (c *Context) DisableVertexAttribArray(index int) {
	c.Call("disableVertexAttribArray", index)
}

// Render geometric primitives from bound and enabled vertex data.
func (c *Context) DrawArrays(mode, first, count int) {
	c.Call("drawArrays", mode, first, count)
}

// Renders geometric primitives indexed by element array data.
func (c *Context) DrawElements(mode, count, typ, offset int) {
	c.Call("drawElements", mode, count, typ, offset)
}

// Turns on specific WebGL capabilities for this context.
func (c *Context) Enable(cap int) {
	c.Call("enable", cap)
}

// Turns on a vertex attribute at a specific index position in
// a vertex attribute array.
func (c *Context) EnableVertexAttribArray(index int) {
	c.Call("enableVertexAttribArray", index)
}

func (c *Context) Finish() {
	c.Call("finish")
}

func (c *Context) Flush() {
	c.Call("flush")
}

// Attaches a WebGLRenderbuffer object as a logical buffer to the
// currently bound WebGLFramebuffer object.
func (c *Context) FrameBufferRenderBuffer(target, attachment, renderbufferTarget int, renderbuffer js.Value) {
	c.Call("framebufferRenderBuffer", target, attachment, renderbufferTarget, renderbuffer)
}

// Attaches a texture to a WebGLFramebuffer object.
func (c *Context) FramebufferTexture2D(target, attachment, textarget int, texture js.Value, level int) {
	c.Call("framebufferTexture2D", target, attachment, textarget, texture, level)
}

// Sets whether or not polygons are considered front-facing based
// on their winding direction.
func (c *Context) FrontFace(mode int) {
	c.Call("frontFace", mode)
}

// Creates a set of textures for a WebGLTexture object with image
// dimensions from the original size of the image down to a 1x1 image.
func (c *Context) GenerateMipmap(target int) {
	c.Call("generateMipmap", target)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a vertex attribute at a specific index position in a program object.
func (c *Context) GetActiveAttrib(program js.Value, index int) js.Value {
	return c.Call("getActiveAttrib", program, index)
}

// Returns an WebGLActiveInfo object containing the size, type, and name
// of a uniform attribute at a specific index position in a program object.
func (c *Context) GetActiveUniform(program js.Value, index int) js.Value {
	return c.Call("getActiveUniform", program, index)
}

// Returns a slice of WebGLShaders bound to a WebGLProgram.
func (c *Context) GetAttachedShaders(program js.Value) []js.Value {
	objs := c.Call("getAttachedShaders", program)
	shaders := make([]js.Value, objs.Length())
	for i := 0; i < objs.Length(); i++ {
		shaders[i] = objs.Index(i)
	}
	return shaders
}

// Returns an index to the location in a program of a named attribute variable.
func (c *Context) GetAttribLocation(program js.Value, name string) int {
	return c.Call("getAttribLocation", program, name).Int()
}

// TODO: Create type specific variations.
// Returns the type of a parameter for a given buffer.
func (c *Context) GetBufferParameter(target, pname int) js.Value {
	return c.Call("getBufferParameter", target, pname)
}

// TODO: Create type specific variations.
// Returns the natural type value for a constant parameter.
func (c *Context) GetParameter(pname int) js.Value {
	return c.Call("getParameter", pname)
}

// Returns a value for the WebGL error flag and clears the flag.
func (c *Context) GetError() int {
	return c.Call("getError").Int()
}

// TODO: Create type specific variations.
// Enables a passed extension, otherwise returns null.
func (c *Context) GetExtension(name string) js.Value {
	return c.Call("getExtension", name)
}

// TODO: Create type specific variations.
// Gets a parameter value for a given target and attachment.
func (c *Context) GetFramebufferAttachmentParameter(target, attachment, pname int) js.Value {
	return c.Call("getFramebufferAttachmentParameter", target, attachment, pname)
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as an int.
func (c *Context) GetProgramParameteri(program js.Value, pname int) int {
	return c.Call("getProgramParameter", program, pname).Int()
}

// Returns the value of the program parameter that corresponds to a supplied pname
// which is interpreted as a bool.
func (c *Context) GetProgramParameterb(program js.Value, pname int) bool {
	return c.Call("getProgramParameter", program, pname).Bool()
}

// Returns information about the last error that occurred during
// the failed linking or validation of a WebGL program object.
func (c *Context) GetProgramInfoLog(program js.Value) string {
	return c.Call("getProgramInfoLog", program).String()
}

// TODO: Create type specific variations.
// Returns a renderbuffer parameter from the currently bound WebGLRenderbuffer object.
func (c *Context) GetRenderbufferParameter(target, pname int) js.Value {
	return c.Call("getRenderbufferParameter", target, pname)
}

// TODO: Create type specific variations.
// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameter(shader js.Value, pname int) js.Value {
	return c.Call("getShaderParameter", shader, pname)
}

// Returns the value of the parameter associated with pname for a shader object.
func (c *Context) GetShaderParameterb(shader js.Value, pname int) bool {
	return c.Call("getShaderParameter", shader, pname).Bool()
}

// Returns errors which occur when compiling a shader.
func (c *Context) GetShaderInfoLog(shader js.Value) string {
	return c.Call("getShaderInfoLog", shader).String()
}

// Returns source code string associated with a shader object.
func (c *Context) GetShaderSource(shader js.Value) string {
	return c.Call("getShaderSource", shader).String()
}

// Returns a slice of supported extension strings.
func (c *Context) GetSupportedExtensions() []string {
	ext := c.Call("getSupportedExtensions")
	extensions := make([]string, ext.Length())
	for i := 0; i < ext.Length(); i++ {
		extensions[i] = ext.Index(i).String()
	}
	return extensions
}

// TODO: Create type specific variations.
// Returns the value for a parameter on an active texture unit.
func (c *Context) GetTexParameter(target, pname int) js.Value {
	return c.Call("getTexParameter", target, pname)
}

// TODO: Create type specific variations.
// Gets the uniform value for a specific location in a program.
func (c *Context) GetUniform(program, location js.Value) js.Value {
	return c.Call("getUniform", program, location)
}

// Returns a WebGLUniformLocation object for the location
// of a uniform variable within a WebGLProgram object.
func (c *Context) GetUniformLocation(program js.Value, name string) js.Value {
	return c.Call("getUniformLocation", program, name)
}

// TODO: Create type specific variations.
// Returns data for a particular characteristic of a vertex
// attribute at an index in a vertex attribute array.
func (c *Context) GetVertexAttrib(index, pname int) js.Value {
	return c.Call("getVertexAttrib", index, pname)
}

// Returns the address of a specified vertex attribute.
func (c *Context) GetVertexAttribOffset(index, pname int) int {
	return c.Call("getVertexAttribOffset", index, pname).Int()
}

// public function hint(target:GLenum, mode:GLenum) : Void;

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsBuffer(buffer js.Value) bool {
	return c.Call("isBuffer", buffer).Bool()
}

// Returns whether the WebGL context has been lost.
func (c *Context) IsContextLost() bool {
	return c.Call("isContextLost").Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsFramebuffer(framebuffer js.Value) bool {
	return c.Call("isFramebuffer", framebuffer).Bool()
}

// Returns true if program object is valid, false otherwise.
func (c *Context) IsProgram(program js.Value) bool {
	return c.Call("isProgram", program).Bool()
}

// Returns true if buffer is valid, false otherwise.
func (c *Context) IsRenderbuffer(renderbuffer js.Value) bool {
	return c.Call("isRenderbuffer", renderbuffer).Bool()
}

// Returns true if shader is valid, false otherwise.
func (c *Context) IsShader(shader js.Value) bool {
	return c.Call("isShader", shader).Bool()
}

// Returns true if texture is valid, false otherwise.
func (c *Context) IsTexture(texture js.Value) bool {
	return c.Call("isTexture", texture).Bool()
}

// Returns whether or not a WebGL capability is enabled for this context.
func (c *Context) IsEnabled(capability int) bool {
	return c.Call("isEnabled", capability).Bool()
}

// Sets the width of lines in WebGL.
func (c *Context) LineWidth(width float64) {
	c.Call("lineWidth", width)
}

// Links an attached vertex shader and an attached fragment shader
// to a program so it can be used by the graphics processing unit (GPU).
func (c *Context) LinkProgram(program js.Value) {
	c.Call("linkProgram", program)
}

// Sets pixel storage modes for readPixels and unpacking of textures
// with texImage2D and texSubImage2D.
func (c *Context) PixelStorei(pname, param int) {
	c.Call("pixelStorei", pname, param)
}

// Sets the implementation-specific units and scale factor
// used to calculate fragment depth values.
func (c *Context) PolygonOffset(factor, units float64) {
	c.Call("polygonOffset", factor, units)
}

// TODO: Figure out if pixels should be a slice.
// Reads pixel data into an ArrayBufferView object from a
// rectangular area in the color buffer of the active frame buffer.
func (c *Context) ReadPixels(x, y, width, height, format, typ int, pixels js.Value) {
	c.Call("readPixels", x, y, width, height, format, typ, pixels)
}

// Creates or replaces the data store for the currently bound WebGLRenderbuffer object.
func (c *Context) RenderbufferStorage(target, internalFormat, width, height int) {
	c.Call("renderbufferStorage", target, internalFormat, width, height)
}

//func (c *Context) SampleCoverage(value float64, invert bool) {
//	c.Call("sampleCoverage", value, invert)
//}

// Sets the dimensions of the scissor box.
func (c *Context) Scissor(x, y, width, height int) {
	c.Call("scissor", x, y, width, height)
}

// Sets and replaces shader source code in a shader object.
func (c *Context) ShaderSource(shader js.Value, source string) {
	c.Call("shaderSource", shader, source)
}

// public function stencilFunc(func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilFuncSeparate(face:GLenum, func:GLenum, ref:GLint, mask:GLuint) : Void;
// public function stencilMask(mask:GLuint) : Void;
// public function stencilMaskSeparate(face:GLenum, mask:GLuint) : Void;
// public function stencilOp(fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;
// public function stencilOpSeparate(face:GLenum, fail:GLenum, zfail:GLenum, zpass:GLenum) : Void;

// Loads the supplied pixel data into a texture.
func (c *Context) TexImage2D(target, level, internalFormat, format, kind int, image js.Value) {
	c.Call("texImage2D", target, level, internalFormat, format, kind, image)
}

// Sets texture parameters for the current texture unit.
func (c *Context) TexParameteri(target int, pname int, param int) {
	c.Call("texParameteri", target, pname, param)
}

// Replaces a portion of an existing 2D texture image with all of another image.
func (c *Context) TexSubImage2D(target, level, xoffset, yoffset, format, typ int, image js.Value) {
	c.Call("texSubImage2D", target, level, xoffset, yoffset, format, typ, image)
}

// Assigns a floating point value to a uniform variable for the current program object.
func (c *Context) Uniform1f(location js.Value, x float32) {
	c.Call("uniform1f", location, x)
}

// Assigns a integer value to a uniform variable for the current program object.
func (c *Context) Uniform1i(location js.Value, x int) {
	c.Call("uniform1i", location, x)
}

// Assigns 2 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform2f(location js.Value, x, y float32) {
	c.Call("uniform2f", location, x, y)
}

// Assigns 2 integer values to a uniform variable for the current program object.
func (c *Context) Uniform2i(location js.Value, x, y int) {
	c.Call("uniform2i", location, x, y)
}

// Assigns 3 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform3f(location js.Value, x, y, z float32) {
	c.Call("uniform3f", location, x, y, z)
}

// Assigns 3 integer values to a uniform variable for the current program object.
func (c *Context) Uniform3i(location js.Value, x, y, z int) {
	c.Call("uniform3i", location, x, y, z)
}

// Assigns 4 floating point values to a uniform variable for the current program object.
func (c *Context) Uniform4f(location js.Value, x, y, z, w float32) {
	c.Call("uniform4f", location, x, y, z, w)
}

// Assigns 4 integer values to a uniform variable for the current program object.
func (c *Context) Uniform4i(location js.Value, x, y, z, w int) {
	c.Call("uniform4i", location, x, y, z, w)
}

// public function uniform1fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform1iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform2fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform2iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform3fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform3iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;
// public function uniform4fv(location:WebGLUniformLocation, v:ArrayAccess<Float>) : Void;
// public function uniform4iv(location:WebGLUniformLocation, v:ArrayAccess<Long>) : Void;

// Sets values for a 2x2 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix2fv(location js.Value, transpose bool, value []float32) {
	c.Call("uniformMatrix2fv", location, transpose, value)
}

// Sets values for a 3x3 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix3fv(location js.Value, transpose bool, value []float32) {
	c.Call("uniformMatrix3fv", location, transpose, value)
}

// Sets values for a 4x4 floating point vector matrix into a
// uniform location as a matrix or a matrix array.
func (c *Context) UniformMatrix4fv(location js.Value, transpose bool, value []float32) {
	c.Call("uniformMatrix4fv", location, transpose, value)
}

// Set the program object to use for rendering.
func (c *Context) UseProgram(program js.Value) {
	c.Call("useProgram", program)
}

// Returns whether a given program can run in the current WebGL state.
func (c *Context) ValidateProgram(program js.Value) {
	c.Call("validateProgram", program)
}

func (c *Context) VertexAttribPointer(index, size, typ int, normal bool, stride int, offset int) {
	c.Call("vertexAttribPointer", index, size, typ, normal, stride, offset)
}

// public function vertexAttrib1f(indx:GLuint, x:GLfloat) : Void;
// public function vertexAttrib2f(indx:GLuint, x:GLfloat, y:GLfloat) : Void;
// public function vertexAttrib3f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat) : Void;
// public function vertexAttrib4f(indx:GLuint, x:GLfloat, y:GLfloat, z:GLfloat, w:GLfloat) : Void;
// public function vertexAttrib1fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib2fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib3fv(indx:GLuint, values:ArrayAccess<Float>) : Void;
// public function vertexAttrib4fv(indx:GLuint, values:ArrayAccess<Float>) : Void;

// Represents a rectangular viewable area that contains
// the rendering results of the drawing buffer.
func (c *Context) Viewport(x, y, width, height int) {
	c.Call("viewport", x, y, width, height)
}
