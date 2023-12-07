import { Card } from "@mui/material";
import { useParams } from "react-router";

export default function ProofParam() {
  const { id } = useParams()

  return (
    <>
    <p>Lorem ipsum, dolor sit amet consectetur adipisicing elit. Assumenda, aspernatur aperiam architecto rerum accusantium non minus repellat tenetur</p> <br/> <br/> <p>doloremque, perferendis harum quidem quas obcaecati recusandae, omnis adipisci excepturi ipsa laboriosam?</p>
    <br />
    <br />
    <br />
    <br />
    <Card>
      <p>Card</p>
      <p>{id}</p>
    </Card>
    </>
  );
}
