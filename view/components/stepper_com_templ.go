// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Stepper() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<!-- Stepper --><ul class=\"relative flex flex-row gap-x-2 mb-8 m-w-120\"><!-- Item --><li class=\"flex items-center gap-x-2 shrink basis-0 flex-1 group\"><div class=\"min-w-[28px] min-h-[28px] inline-flex justify-center items-center text-xs align-middle\"><span class=\"w-7 h-7 flex justify-center items-center flex-shrink-0 bg-gray-800 font-medium text-white rounded-full dark:bg-white dark:text-gray-800\">1</span> <span class=\"ms-2 block text-sm font-medium text-gray-800 dark:text-white\">Personal Information</span></div><div class=\"w-full h-px flex-1 bg-gray-200 group-last:hidden dark:bg-gray-700\"></div></li><!-- End Item --><!-- Item --><li class=\"flex items-center gap-x-2 shrink basis-0 flex-1 group\"><div class=\"min-w-[28px] min-h-[28px] inline-flex justify-center items-center text-xs align-middle\"><span class=\"w-7 h-7 flex justify-center items-center flex-shrink-0 bg-white border border-gray-200 font-medium text-gray-800 rounded-full dark:bg-white dark:text-gray-800\">2</span> <span class=\"ms-2 block text-sm font-medium text-gray-800 dark:text-white\">Phone Number Verification</span></div><div class=\"w-full h-px flex-1 bg-gray-200 group-last:hidden dark:bg-gray-700\"></div></li><!-- End Item --></ul><!-- End Stepper -->")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
