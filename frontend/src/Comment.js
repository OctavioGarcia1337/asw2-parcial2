import CommentForm from "./CommentForm";
import "./css/Item.css";

const Comment = ({
  comment,
  activeComment,
  firstName,
  item_id,

}) => {
  const isEditing =
    activeComment &&
    activeComment.id === comment.id &&
    activeComment.type === "editing";
  const createdAt = new Date(comment.createdAt).toLocaleDateString();

  return (
    <div key={comment.id} className="comment">
      <div className="comment-image-container">
        <img src="/user-icon.png" />
      </div>
      <div className="comment-right-part">
        <div className="comment-content">
          <div className="comment-author">{firstName}</div>
          <div>{createdAt}</div>
        </div>
        {!isEditing && <div className="comment-text">{comment.body}</div>}
        <div className="comment-actions">

        </div>

      </div>
    </div>
  );
};

export default Comment;
