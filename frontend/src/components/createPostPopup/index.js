import {useRef, useState} from "react";
import "./style.css";
import EmojiPickerBackgrounds from "./EmojiPickerBackgrounds";
import AddToYourPost from "./AddToYourPost";
import ImagePreview from "./ImagePreview";
import {useClickOutside} from "../../hooks/useClickOutside";
import {createPost} from "../../functions/post";
import PulseLoader from "react-spinners/PulseLoader";
import PostError from "./PostError";
import {uploadImage} from "../../functions/uploadImage";
import dataURItoBlob from "../../helpers/dataURItoBlob";

export default function CreatePostPopup({user, setVisible}) {
    const popup = useRef(null);
    const [text, setText] = useState("");
    const [showPrev, setShowPrev] = useState(false);

    const [loading, setLoading] = useState(false);
    const [error, setError] = useState("");

    const [images, setImages] = useState([]);
    const [background, setBackground] = useState("");
    useClickOutside(popup, () => {
        setVisible(false);
    });

    async function createFbPost(type, background, text, imageURLS) {
        const response = await createPost(type, background, text, imageURLS)
        if (response.success) {
            setLoading(false)
            setText('')
            setImages([])
            setVisible(false)
        } else {
            setError(response.success)
        }
        return response
    }

    const postSubmit = async () => {
        console.log('Click')
        if (background) {
            await createFbPost(null, background, text, null)
        } else if (images && images.length) {
            setLoading(true)
            const uploadImagesResponse = await uploadImage(images.map(img => dataURItoBlob(img)));
            if (!uploadImagesResponse.success) {
                setLoading(false)
                setError(uploadImagesResponse.message)
                return
            }
            if (!uploadImagesResponse.images) {
                setLoading(false)
                setError('Upload images failed')
                return
            }
            const createPostResponse = await createFbPost(null, null, text, uploadImagesResponse.images);
            if (!createPostResponse.success) {
                setLoading(false)
                setError(createPostResponse.message)
                return
            }
            setLoading(false)
            setText('')
            setImages([])
            setVisible(false)
        } else if (text) {
            await createFbPost(null, background, text, null)
        } else {
            console.log("Nothing here")
        }
        setImages([])
        setText('')
        setLoading(false)
    };


    return (
        <div className="blur">
            <div className="postBox" ref={popup}>
                {error && <PostError error={error} setError={setError}/>}
                <div className="box_header">
                    <div
                        className="small_circle"
                        onClick={() => {
                            setVisible(false);
                        }}
                    >
                        <i className="exit_icon"></i>
                    </div>
                    <span>Create Post</span>
                </div>
                <div className="box_profile">
                    <img src={user.picture} alt="" className="box_profile_img"/>
                    <div className="box_col">
                        <div className="box_profile_name">
                            {user.first_name} {user.last_name}
                        </div>
                        <div className="box_privacy">
                            <img src="../../../icons/public.png" alt=""/>
                            <span>Public</span>
                            <i className="arrowDown_icon"></i>
                        </div>
                    </div>
                </div>

                {!showPrev ? (
                    <>
                        <EmojiPickerBackgrounds
                            text={text}
                            user={user}
                            setText={setText}
                            showPrev={showPrev}
                            setBackground={setBackground}
                            background={background}
                        />
                    </>
                ) : (
                    <ImagePreview
                        text={text}
                        user={user}
                        setText={setText}
                        showPrev={showPrev}
                        images={images}
                        setImages={setImages}
                        setShowPrev={setShowPrev}
                    />
                )}
                <AddToYourPost setShowPrev={setShowPrev}/>
                <button
                    className="post_submit"
                    onClick={() => {
                        postSubmit();
                    }}
                    disabled={loading}
                >
                    {loading ? <PulseLoader color="#fff" size={5}/> : "Post"}
                </button>
            </div>
        </div>
    );
}
