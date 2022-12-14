import { useState, useEffect } from "react";
import CommentForm from "./CommentForm";
import Comment from "./Comment";
import {
  getComments as getCommentsApi,
  createComment as createCommentApi,
} from "./CommentBD";
import "./css/Item.css";




const Comments = ({ first_name, uid, item}) => {
  const [backendComments, setBackendComments] = useState([]);
  const [activeComment, setActiveComment] = useState(null);
  const rootComments = backendComments.filter((backendComment) => backendComment.item_id === item);

  const addComment = (text) => {
    createCommentApi(text, uid, item).then((comment) => {
      setBackendComments([comment, ...backendComments]);
      setActiveComment(null);
    });
  };

  useEffect(() => {
    getCommentsApi().then((data) => {
      setBackendComments(data);
    });
  }, []);

  return (
    <div className="comments">
      <h3 className="comments-title">Comments</h3>
      <div className="comment-form-title">Write comment</div>
      <CommentForm submitLabel="Write" handleSubmit={addComment} />
      <div className="comments-container">
        {rootComments.map((rootComment) => (
          <Comment
            key={rootComment.id}
            comment={rootComment}
            activeComment={activeComment}
            setActiveComment={setActiveComment}
            addComment={addComment}
            firstName={first_name}
          />
        ))}
      </div>
    </div>
  );
};

export default Comments;
