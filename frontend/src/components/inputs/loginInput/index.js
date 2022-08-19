import {ErrorMessage, useField} from "formik";
import "./style.css"
import classNames from "classnames";
import {useMediaQuery} from "react-responsive";

export default function Index({placeholder, bottom, ...props}) {
    const [field, meta] = useField(props);
    const isDesktop = useMediaQuery({query: '(min-width: 850px)'})
    console.log(isDesktop)
    return (
        <div className="input_wrap">
            {meta.touched && meta.error && !bottom && (
                <div className={classNames({
                    "input_error_desktop input_error": isDesktop,
                    "input_error": !isDesktop
                })}
                     style={{transform: "translateY(3px)"}}>
                    {meta.touched && meta.error && <ErrorMessage name={field.name}/>}
                    {meta.touched && meta.error && (
                        <div className={classNames({
                            'error_arrow_left': isDesktop,
                            "error_arrow_top": !isDesktop
                        })}></div>
                    )}
                </div>
            )}
            <input
                className={classNames({input_error_border: meta.touched && meta.error})}
                type={field.type}
                name={field.name}
                placeholder={placeholder}
                {...field}
                {...props}
            />
            {meta.touched && meta.error && bottom && (
                <div className={classNames({
                    "input_error_desktop input_error": isDesktop,
                    "input_error": !isDesktop
                })} style={{transform: "translateY(2px)"}}>
                    {meta.touched && meta.error && <ErrorMessage name={field.name}/>}
                    {meta.touched && meta.error && (
                        <div className={classNames({
                            "error_arrow_left": isDesktop,
                            "error_arrow_bottom": !isDesktop
                        })}></div>
                    )}
                </div>
            )}

            {meta.touched && meta.error && (
                <i className="error_icon" style={{top: `${!bottom && "63%"}`}}></i>
            )}
        </div>

    );
}
